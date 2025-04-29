import api from './index';

export const fetchNfts = async () => {
  try {
    const response = await api.get('/nfts');
    return response.data;
  } catch (error) {
    console.error('Error fetching NFTs:', error);
    throw error; // 에러를 호출한 쪽으로 전달
  }
};