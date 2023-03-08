import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL;

export function getFizzBuzz(count: number) {
  return axios.post(`${API_URL}/fizzbuzz`, {count: count}).then((response) => response.data);
}
