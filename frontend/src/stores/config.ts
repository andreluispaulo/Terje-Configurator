import { defineStore } from 'pinia';
import { ref } from 'vue';
import { getFile, saveFile as apiSaveFile } from '../services/api';
import type { FileContent } from '../types';

export const useConfigStore = defineStore('config', () => {
  const currentPath = ref<string | null>(null);
  const currentContent = ref<FileContent | null>(null);
  const isDirty = ref(false);
  const loading = ref(false);
  const error = ref<string | null>(null);

  async function loadFile(path: string) {
    loading.value = true;
    error.value = null;
    try {
      const response = await getFile(path);
      currentContent.value = response.data;
      currentPath.value = path;
      isDirty.value = false;
    } catch (err: any) {
      error.value = err.message;
    } finally {
      loading.value = false;
    }
  }

  async function saveChanges(updates: any[]) {
    if (!currentPath.value) return;
    loading.value = true;
    try {
      await apiSaveFile(currentPath.value, updates);
      isDirty.value = false;
      await loadFile(currentPath.value);
    } catch (err: any) {
      error.value = err.message;
    } finally {
      loading.value = false;
    }
  }

  return {
    currentPath,
    currentContent,
    isDirty,
    loading,
    error,
    loadFile,
    saveChanges
  };
});
