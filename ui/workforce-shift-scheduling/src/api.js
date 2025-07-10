import axios from 'axios';

const API_BASE = 'http://localhost:8080';

export const getSchedule = async (payload) => {
  const response = await axios.post(`${API_BASE}/schedule`, payload);
  return response.data;
};