import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://127.0.0.1:3001',
});

// instance.defaults.headers.common['Authorization'] = 'AUTH TOKEN';
instance.defaults.headers.post['Content-Type'] = 'application/json';
instance.defaults.withCredentials = true

export default instance;