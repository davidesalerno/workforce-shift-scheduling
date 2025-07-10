import axios from 'axios';

const API_BASE = 'http://localhost:8080';

export const getSchedule = async (payload) => {
  const apiUrl = import.meta.env.VITE_BASE_URL_OVERRIDE || API_BASE;
  console.log(import.meta.env.VITE_BASE_URL_OVERRIDE) 
  console.log(apiUrl); 
  const response = await axios.post(`${apiUrl}/schedule`, payload);
  return response.data;
};