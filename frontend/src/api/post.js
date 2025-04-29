import api from './index';

// 게시물 목록 가져오기
export const fetchPosts = async () => {
  const response = await api.get('/posts');
  return response.data;
};

// 게시물 생성
export const createPost = async (postData) => {
  const response = await api.post('/posts', postData);
  return response.data;
};

// 게시물 삭제
export const deletePost = async (postId) => {
  const response = await api.delete(`/posts/${postId}`);
  return response.data;
};