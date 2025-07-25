package service

import (
	"github.com/google/uuid"
	"github.com/ursulgwopp/pulse-api/internal/entity"
	"github.com/ursulgwopp/pulse-api/internal/errors"
)

func (s *Service) NewPost(login string, req entity.NewPostRequest) (entity.Post, error) {
	if err := validateContent(req.Content); err != nil {
		return entity.Post{}, err
	}

	for _, tag := range req.Tags {
		if !isValidTag(tag) {
			return entity.Post{}, errors.ErrInvalidTag
		}
	}

	return s.repo.NewPost(login, req)
}

func (s *Service) GetPost(login string, postId uuid.UUID) (entity.Post, error) {
	// id exists
	exists, err := s.repo.CheckPostIdExists(postId)
	if err != nil {
		return entity.Post{}, err
	}

	if !exists {
		return entity.Post{}, errors.ErrPostIdNotFound
	}

	post, err := s.repo.GetPost(postId)
	if err != nil {
		return entity.Post{}, err
	}

	// my posts
	if login == post.Author {
		return post, nil
	}

	// public profile
	is_public, err := s.repo.CheckProfileIsPublic(post.Author)
	if err != nil {
		return entity.Post{}, err
	}

	if is_public {
		return post, nil
	}

	friends, err := s.repo.ListFriends(post.Author, 1000000, 0)
	if err != nil {
		return entity.Post{}, err
	}

	if isFriend(friends, login) {
		return post, nil
	}

	return entity.Post{}, errors.ErrAccessDenied
}

func (s *Service) ListMyPosts(login string, limit int, offset int) ([]entity.Post, error) {
	if limit < 0 || offset < 0 {
		return []entity.Post{}, errors.ErrInvalidPaginationParams
	}

	return s.repo.ListPosts(login, limit, offset)
}

func (s *Service) ListPosts(userLogin string, login string, limit int, offset int) ([]entity.Post, error) {
	// pagination params
	if limit < 0 || offset < 0 {
		return []entity.Post{}, errors.ErrInvalidPaginationParams
	}

	// login exists
	exists, err := s.repo.CheckLoginExists(login)
	if err != nil {
		return []entity.Post{}, err
	}

	if !exists {
		return []entity.Post{}, errors.ErrLoginDoesNotExist
	}

	posts, err := s.repo.ListPosts(login, limit, offset)
	if err != nil {
		return []entity.Post{}, err
	}

	// my posts
	if userLogin == login {
		return posts, nil
	}

	// public profile
	is_public, err := s.repo.CheckProfileIsPublic(login)
	if err != nil {
		return []entity.Post{}, err
	}

	if is_public {
		return posts, nil
	}

	// follows
	friends, err := s.repo.ListFriends(login, 1000000, 0)
	if err != nil {
		return []entity.Post{}, err
	}

	if isFriend(friends, userLogin) {
		return posts, nil
	}

	return []entity.Post{}, errors.ErrAccessDenied
}

func (s *Service) LikePost(login string, postId uuid.UUID) (entity.Post, error) {
	// id exists
	exists, err := s.repo.CheckPostIdExists(postId)
	if err != nil {
		return entity.Post{}, err
	}

	if !exists {
		return entity.Post{}, errors.ErrPostIdNotFound
	}

	author, err := s.repo.CheckPostAuthor(postId)
	if err != nil {
		return entity.Post{}, err
	}

	// public profile
	is_public, err := s.repo.CheckProfileIsPublic(author)
	if err != nil {
		return entity.Post{}, err
	}

	if !is_public && login != author {
		// follows
		friends, err := s.repo.ListFriends(author, 1000000, 0)
		if err != nil {
			return entity.Post{}, err
		}

		if !isFriend(friends, login) {
			return entity.Post{}, errors.ErrAccessDenied
		}
	}

	post, err := s.repo.LikePost(login, postId)
	if err != nil {
		return entity.Post{}, err
	}

	return post, nil
}

func (s *Service) DislikePost(login string, postId uuid.UUID) (entity.Post, error) {
	// id exists
	exists, err := s.repo.CheckPostIdExists(postId)
	if err != nil {
		return entity.Post{}, err
	}

	if !exists {
		return entity.Post{}, errors.ErrPostIdNotFound
	}

	author, err := s.repo.CheckPostAuthor(postId)
	if err != nil {
		return entity.Post{}, err
	}

	// public profile
	is_public, err := s.repo.CheckProfileIsPublic(author)
	if err != nil {
		return entity.Post{}, err
	}

	if !is_public && login != author {
		// follows
		friends, err := s.repo.ListFriends(author, 1000000, 0)
		if err != nil {
			return entity.Post{}, err
		}

		if !isFriend(friends, login) {
			return entity.Post{}, errors.ErrAccessDenied
		}
	}

	post, err := s.repo.DislikePost(login, postId)
	if err != nil {
		return entity.Post{}, err
	}

	return post, nil
}
