import axios from "axios";

export const api = axios.create({
    baseURL: "/api"
})
export const ws = new WebSocket("/new-load-added");
