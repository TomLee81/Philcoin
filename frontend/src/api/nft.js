import api from './index';

export const fetchNfts = async () => {
  const response = await api.get('/nfts');
  return response.data;
};