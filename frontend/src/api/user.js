import api from './index';

// 사용자 정보 가져오기
export const fetchUser = async (userId) => {
  const response = await api.get(`/users/${userId}`);
  return response.data;
};

// 사용자 정보 업데이트
export const updateUser = async (userId, userData) => {
  const response = await api.put(`/users/${userId}`, userData);
  return response.data;
};

// 사용자 삭제
export const deleteUser = async (userId) => {
  const response = await api.delete(`/users/${userId}`);
  return response.data;
};