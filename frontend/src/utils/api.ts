import axios from "axios";

const apiUrl = import.meta.env.APIURL;

export const API = axios.create({
    baseURL: apiUrl,
    withCredentials: true,
});

export default apiUrl;