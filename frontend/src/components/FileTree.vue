<script setup lang="ts">
import { ref, computed } from 'vue';
import { Folder, FileText, ChevronRight, ChevronDown, FolderOpen } from 'lucide-vue-next';
import type { TreeNode } from '../services/api';
import { useConfigStore } from '../stores/config';

const props = defineProps<{
  node: TreeNode;
  level?: number;
}>();

const store = useConfigStore();
const isOpen = ref(false);
const isFolder = computed(() => props.node.type === 'folder');
const paddingLeft = computed(() => `${(props.level || 0) * 1.25}rem`);

function toggle() {
  if (isFolder.value) {
    isOpen.value = !isOpen.value;
  } else {
    store.loadFile(props.node.path);
  }
}
</script>

<template>
  <div class="select-none">
    <div
      class="flex items-center py-1.5 px-2 rounded-md cursor-pointer transition-all duration-200 group relative overflow-hidden"
      :class="{ 
        'bg-primary/10 text-primary font-medium': !isFolder && store.currentPath === node.path,
        'text-muted-foreground hover:text-foreground hover:bg-muted/50': !(!isFolder && store.currentPath === node.path)
      }"
      :style="{ paddingLeft }"
      @click="toggle"
    >
      <!-- Active Indicator Line -->
      <div 
        v-if="!isFolder && store.currentPath === node.path" 
        class="absolute left-0 top-1/2 -translate-y-1/2 w-0.5 h-4 bg-primary rounded-r-full"
      ></div>

      <!-- Icon Area -->
      <div class="mr-2 flex items-center justify-center w-5 h-5 shrink-0">
        <template v-if="isFolder">
            <ChevronDown v-if="isOpen" :size="14" class="text-muted-foreground/70" />
            <ChevronRight v-else :size="14" class="text-muted-foreground/70" />
        </template>
      </div>
      
      <div class="mr-2 shrink-0">
         <FolderOpen v-if="isFolder && isOpen" :size="16" class="text-amber-500/90" />
         <Folder v-else-if="isFolder" :size="16" class="text-amber-500/70 group-hover:text-amber-500 transition-colors" />
         <FileText v-else :size="16" class="transition-colors" :class="store.currentPath === node.path ? 'text-primary' : 'text-blue-400/60 group-hover:text-blue-400'" />
      </div>
      
      <span class="text-sm truncate leading-none pt-0.5">{{ node.name }}</span>
    </div>

    <div v-if="isFolder && isOpen" class="relative">
      <!-- Indentation Guide Line -->
      <div 
        class="absolute left-[1.2rem] top-0 bottom-0 w-px bg-border/40 ml-[calc(var(--level)*1.25rem)]"
        :style="{ left: `${(level || 0) * 1.25 + 0.9}rem` }"
      ></div>
      
      <FileTree
        v-for="child in node.children"
        :key="child.path"
        :node="child"
        :level="(level || 0) + 1"
      />
    </div>
  </div>
</template>
