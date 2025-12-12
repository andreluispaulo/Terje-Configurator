<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { getTree, type TreeNode } from '../services/api';
import FileTree from './FileTree.vue';
import { Settings2, Languages } from 'lucide-vue-next';
import { useI18nStore } from '../stores/i18n';

const tree = ref<TreeNode[]>([]);
const i18n = useI18nStore();

onMounted(async () => {
  try {
    const res = await getTree();
    tree.value = res.data;
  } catch (e) {
    console.error(e);
  }
});
</script>

<template>
  <div class="h-full bg-card border-r border-border flex flex-col w-72 transition-all duration-300">
    <!-- Header -->
    <div class="p-6 border-b border-border flex items-center gap-3 shrink-0">
      <div class="h-10 w-10 rounded-xl bg-primary/10 flex items-center justify-center text-primary ring-1 ring-primary/20 shadow-sm">
        <Settings2 :size="20" stroke-width="2.5" />
      </div>
      <div>
        <h2 class="font-bold text-foreground tracking-tight leading-none">{{ i18n.t.title }}</h2>
        <span class="text-xs text-muted-foreground font-medium">Configurator v1.0</span>
      </div>
    </div>

    <!-- Tree -->
    <div class="flex-1 py-4 px-3 overflow-y-auto scrollbar-thin scrollbar-thumb-muted scrollbar-track-transparent">
      <div class="space-y-0.5">
        <FileTree v-for="node in tree" :key="node.path" :node="node" />
      </div>
    </div>

    <!-- Footer / Language -->
    <div class="p-4 border-t border-border bg-muted/20 shrink-0">
      <div class="flex items-center gap-2 text-muted-foreground mb-3 px-1">
        <Languages :size="14" />
        <span class="text-[10px] font-bold uppercase tracking-widest opacity-70">Region</span>
      </div>
      <div class="grid grid-cols-4 gap-1 p-1 bg-background rounded-lg border border-input shadow-sm">
        <button 
          @click="i18n.setLocale('pt-BR')"
          class="px-2 py-1.5 text-xs rounded-md transition-all font-medium flex items-center justify-center"
          :class="i18n.currentLocale === 'pt-BR' ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground hover:bg-muted'"
        >PT</button>
        <button 
          @click="i18n.setLocale('en-US')"
          class="px-2 py-1.5 text-xs rounded-md transition-all font-medium flex items-center justify-center"
          :class="i18n.currentLocale === 'en-US' ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground hover:bg-muted'"
        >EN</button>
        <button 
          @click="i18n.setLocale('es-ES')"
          class="px-2 py-1.5 text-xs rounded-md transition-all font-medium flex items-center justify-center"
          :class="i18n.currentLocale === 'es-ES' ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground hover:bg-muted'"
        >ES</button>
        <button 
          @click="i18n.setLocale('ru-RU')"
          class="px-2 py-1.5 text-xs rounded-md transition-all font-medium flex items-center justify-center"
          :class="i18n.currentLocale === 'ru-RU' ? 'bg-primary text-primary-foreground shadow-sm' : 'text-muted-foreground hover:text-foreground hover:bg-muted'"
        >RU</button>
      </div>
    </div>
  </div>
</template>
