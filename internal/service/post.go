package service

import (
	"context"

	"github.com/realdanielursul/pulse-api/internal/entity"
	"github.com/realdanielursul/pulse-api/internal/repository"
	"github.com/realdanielursul/pulse-api/pkg/hasher"
)

type PostService struct {
	postRepo       repository.Post
	userRepo       repository.User
	friendRepo     repository.Friend
	passwordHasher hasher.PasswordHasher
}

func NewPostService(postRepo repository.Post, userRepo repository.User, friendRepo repository.Friend) *PostService {
	return &PostService{
		postRepo:   postRepo,
		userRepo:   userRepo,
		friendRepo: friendRepo,
	}
}

func (s *PostService) CreatePost(ctx context.Context, input *PostCreatePostInput) (*PostOutput, error) {
	postInfo := &entity.Post{
		Content: input.Content,
		Author:  input.Author,
		Tags:    input.Tags,
	}

	post, err := s.postRepo.CreatePost(ctx, postInfo)
	if err != nil {
		return nil, err
	}

	return &PostOutput{
		Id:            post.Id,
		Content:       post.Content,
		Author:        post.Author,
		Tags:          post.Tags,
		CreatedAt:     post.CreatedAt,
		LikesCount:    post.LikesCount,
		DislikesCount: post.DislikesCount,
	}, nil
}

func (s *PostService) GetPost(ctx context.Context, postId, requesterLogin string) (*PostOutput, error) {
	post, err := s.postRepo.GetPostById(ctx, postId)
	if err != nil {
		if post == nil {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	postOutput := &PostOutput{
		Id:            post.Id,
		Content:       post.Content,
		Author:        post.Author,
		Tags:          post.Tags,
		CreatedAt:     post.CreatedAt,
		LikesCount:    post.LikesCount,
		DislikesCount: post.DislikesCount,
	}

	if requesterLogin == post.Author {
		return postOutput, nil
	}

	user, err := s.userRepo.GetUserByLogin(ctx, post.Author)
	if err != nil {
		return nil, err
	}

	if user.IsPublic {
		return postOutput, nil
	}

	isFriend, err := s.friendRepo.IsFriend(ctx, requesterLogin, post.Author)
	if err != nil {
		return nil, err
	}

	if isFriend {
		return postOutput, nil
	}

	return nil, ErrAccessDenied
}

func (s *PostService) GetMyFeed(ctx context.Context, userLogin string, limit, offset int) ([]*PostOutput, error) {
	posts, err := s.postRepo.GetUserPosts(ctx, userLogin, limit, offset)
	if err != nil {
		return nil, err
	}

	postsOutput := make([]*PostOutput, 0, len(posts))
	for _, post := range posts {
		postOutput := &PostOutput{
			Id:            post.Id,
			Content:       post.Content,
			Author:        post.Author,
			Tags:          post.Tags,
			CreatedAt:     post.CreatedAt,
			LikesCount:    post.LikesCount,
			DislikesCount: post.DislikesCount,
		}

		postsOutput = append(postsOutput, postOutput)
	}

	return postsOutput, nil
}

func (s *PostService) GetUserFeed(ctx context.Context, login, requesterLogin string, limit, offset int) ([]*PostOutput, error) {
	posts, err := s.postRepo.GetUserPosts(ctx, login, limit, offset)
	if err != nil {
		return nil, err
	}

	postsOutput := make([]*PostOutput, 0, len(posts))
	for _, post := range posts {
		postOutput := &PostOutput{
			Id:            post.Id,
			Content:       post.Content,
			Author:        post.Author,
			Tags:          post.Tags,
			CreatedAt:     post.CreatedAt,
			LikesCount:    post.LikesCount,
			DislikesCount: post.DislikesCount,
		}

		postsOutput = append(postsOutput, postOutput)
	}

	if login == requesterLogin {
		return postsOutput, nil
	}

	user, err := s.userRepo.GetUserByLogin(ctx, login)
	if err != nil {
		return nil, err
	}

	if user.IsPublic {
		return postsOutput, nil
	}

	isFriend, err := s.friendRepo.IsFriend(ctx, requesterLogin, login)
	if err != nil {
		return nil, err
	}

	if isFriend {
		return postsOutput, nil
	}

	return nil, ErrAccessDenied
}

func (s *PostService) LikePost(ctx context.Context, postId, userLogin string) (*PostOutput, error) {
	post, err := s.postRepo.GetPostById(ctx, postId)
	if err != nil {
		if post == nil {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	if err := s.postRepo.LikePost(ctx, postId, userLogin); err != nil {
		return nil, err
	}

	return s.GetPost(ctx, postId, userLogin)
}

func (s *PostService) DislikePost(ctx context.Context, postId, userLogin string) (*PostOutput, error) {
	post, err := s.postRepo.GetPostById(ctx, postId)
	if err != nil {
		if post == nil {
			return nil, ErrPostNotFound
		}

		return nil, err
	}

	if err := s.postRepo.DislikePost(ctx, postId, userLogin); err != nil {
		return nil, err
	}

	return s.GetPost(ctx, postId, userLogin)
}
