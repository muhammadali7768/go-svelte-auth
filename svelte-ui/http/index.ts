import axios from 'axios';

import Cookies from 'js-cookie';
const API_URL = "http://localhost:8888/api";

const api = axios.create({
  validateStatus: function (status) {
    return status == 200 || status == 201;
  },
  withCredentials: true, // so that cookies are automatically attached to each request
  baseURL: `${API_URL}`,
});

// $api.interceptors.request.use(async config => {
//   console.log("api interceptors",Cookies.get("session-id"));
//   config.headers.Authorization = `Bearer ${Cookies.get('session-id')}`;


//   return config;
// });


// $api.interceptors.response.use(
//   config => {
//     return config;
//   },
//   async error => {
//     const originalRequest = error.config;
//     if (error?.response?.status == 401 && error.config && !error.config._isRetry) {
//       originalRequest._isRetry = true;
//       try {
//         const response = await axios.get<AuthResponse>(`${API_URL}/auth/refresh`, {
//           withCredentials: true,
//           headers: {
//             fingerprint: await fingerprint,
//           },
//         });
//         console.log("refresh api res",response)
//         Cookies.remove('token-access');
//         Cookies.set('token-access', response.data.user.accessToken);
    
//         return $api.request(originalRequest);
//       } catch (e) {
//         console.log({ e });
//         Cookies.remove('token-access');
//         console.log('NOT AUTHORIZED OR INCORRECT ACCESS TOKEN FORMAT!');
//       }
//     }
//     throw error;
//   },
// );

export default api;
