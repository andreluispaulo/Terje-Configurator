import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
});

export const getTree = () => api.get<TreeNode[]>('/tree');
export const getFile = (path: string) => api.get(`/file?path=${encodeURIComponent(path)}`);
export const saveFile = (path: string, updates: any[]) => api.post('/file', { path, updates });
export const getHistory = (path: string) => api.get(`/history?path=${encodeURIComponent(path)}`);
export const restoreFile = (id: number) => api.post('/restore', { id });

export interface TreeNode {
  name: string;
  path: string;
  type: 'file' | 'folder';
  children?: TreeNode[];
}
