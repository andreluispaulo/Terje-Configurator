<script setup lang="ts">
import Sidebar from './components/Sidebar.vue';
import ConfigEditor from './components/ConfigEditor.vue';
</script>

<template>
  <div class="flex h-screen w-screen bg-background text-foreground overflow-hidden font-sans antialiased selection:bg-primary/20 selection:text-primary">
    <Sidebar />
    <main class="flex-1 h-full overflow-hidden relative">
      <ConfigEditor />
    </main>
  </div>
</template>
