<script setup lang="ts">
import { computed, ref, watch } from 'vue';
import { useConfigStore } from '../stores/config';
import { useI18nStore } from '../stores/i18n';
import { Save, Info, Check, X, AlertTriangle, AlertCircle } from 'lucide-vue-next';
import type { CFGFile, XMLFile, CFGLine, XMLLine } from '../types';
import { configDescriptions } from '../locales/configDescriptions';

const store = useConfigStore();
const i18n = useI18nStore();
const activeHelp = ref<number | null>(null);

function getDescription(key: string, originalDescription: string) {
  const currentLocale = i18n.currentLocale;
  const translation = configDescriptions[key];
  
  if (translation && translation[currentLocale]) {
    return translation[currentLocale];
  }
  
  return originalDescription;
}

// Reset updates when path changes
watch(() => store.currentPath, () => {
  updates.value = [];
  activeHelp.value = null;
});

const isCFG = computed(() => store.currentPath?.endsWith('.cfg'));
const isXML = computed(() => store.currentPath?.endsWith('.xml'));

const cfgLines = computed(() => {
  if (isCFG.value && store.currentContent) {
    return (store.currentContent as CFGFile).lines.filter(l => l.type === 1);
  }
  return [];
});

const xmlLines = computed(() => {
  if (isXML.value && store.currentContent) {
    return (store.currentContent as XMLFile).lines.filter(l => l.segments.some(s => s.isAttribute));
  }
  return [];
});

const updates = ref<any[]>([]);

function getDisplayValue(lineIndex: number, key: string, original: string) {
  const update = updates.value.find(u => u.lineIndex === lineIndex && u.key === key);
  return update ? update.value : original;
}

function onUpdateCFG(line: CFGLine, newValue: any) {
  let val = String(newValue);
  const idx = updates.value.findIndex(u => u.lineIndex === line.index);
  if (idx >= 0) {
    updates.value[idx].value = val;
  } else {
    updates.value.push({ lineIndex: line.index, key: line.key, value: val });
  }
  store.isDirty = true;
}

function onUpdateXML(line: XMLLine, attrName: string, newValue: any) {
  let val = String(newValue);
  const idx = updates.value.findIndex(u => u.lineIndex === line.index && u.key === attrName);
  if (idx >= 0) {
    updates.value[idx].value = val;
  } else {
    updates.value.push({ lineIndex: line.index, key: attrName, value: val });
  }
  store.isDirty = true;
}

async function save() {
  await store.saveChanges(updates.value);
  updates.value = [];
}

function toggleHelp(index: number) {
  activeHelp.value = activeHelp.value === index ? null : index;
}

function getInputType(val: string, metadata?: any) {
    // Priority 1: Trust Metadata if available
    if (metadata?.type) {
        const type = metadata.type.toLowerCase().trim();
        if (type === 'bool' || type === 'boolean') return 'bool';
        if (type === 'int' || type === 'integer' || type === 'float' || type === 'number' || type === 'decimal') return 'number';
        return 'text';
    }

    // Priority 2: Value Heuristics (Fallback)
    const lowerVal = val.toLowerCase().trim();
    if (lowerVal === 'true' || lowerVal === 'false') return 'bool';
    if (!isNaN(Number(val)) && val.trim() !== '') return 'number';
    return 'text';
}
</script>

<template>
  <div class="flex flex-col h-full bg-background font-sans relative">
    
    <!-- Top Bar / Header -->
    <div class="h-16 border-b border-border flex items-center justify-between px-6 bg-card/50 backdrop-blur-md sticky top-0 z-30">
      <div class="flex flex-col justify-center">
         <div class="flex items-center gap-2 text-sm text-muted-foreground font-medium">
            <span>Config</span>
            <span class="text-muted-foreground/40">/</span>
            <span v-if="store.currentPath" class="text-foreground">{{ store.currentPath.split('/').pop() }}</span>
            <span v-else class="text-foreground">...</span>
         </div>
      </div>
      
      <button 
        v-if="store.currentPath"
        @click="save"
        :disabled="!store.isDirty || store.loading"
        class="flex items-center gap-2 px-4 py-2 rounded-md transition-all duration-200 text-sm font-medium border"
        :class="(!store.isDirty || store.loading) 
          ? 'bg-muted text-muted-foreground border-transparent cursor-not-allowed opacity-70' 
          : 'bg-primary text-primary-foreground border-primary hover:bg-primary/90 shadow-md shadow-primary/20'"
      >
        <div class="relative">
             <Save :size="16" :class="{ 'opacity-0': store.loading }" />
             <div v-if="store.loading" class="absolute inset-0 animate-spin border-2 border-current border-t-transparent rounded-full"></div>
        </div>
        <span>{{ store.loading ? i18n.t.saving : i18n.t.saveChanges }}</span>
      </button>
    </div>

    <!-- Main Scrollable Area -->
    <div class="flex-1 overflow-y-auto p-6 scrollbar-thin scrollbar-thumb-border scrollbar-track-transparent">
      
      <!-- Loading State -->
      <div v-if="store.loading && !store.currentContent" class="flex flex-col justify-center items-center h-full text-muted-foreground gap-4">
        <div class="w-8 h-8 border-4 border-primary border-t-transparent rounded-full animate-spin"></div>
        <span class="text-sm font-medium animate-pulse">Loading configuration...</span>
      </div>
      
      <!-- Empty State -->
      <div v-else-if="!store.currentPath" class="h-full flex flex-col items-center justify-center text-muted-foreground/60">
        <div class="w-32 h-32 rounded-3xl bg-card border border-border flex items-center justify-center shadow-xl mb-6">
           <FileJson :size="64" stroke-width="1.5" />
        </div>
        <h3 class="text-xl font-semibold text-foreground tracking-tight">{{ i18n.t.selectModule }}</h3>
        <p class="text-sm mt-2 max-w-xs text-center">{{ i18n.t.chooseFile }}</p>
      </div>

      <!-- Content Grid -->
      <div v-else class="max-w-4xl mx-auto space-y-px bg-border rounded-lg overflow-hidden shadow-sm border border-border">
        
        <!-- CFG Content -->
        <template v-if="isCFG">
          <div 
            v-for="line in cfgLines" 
            :key="line.index" 
            class="bg-card hover:bg-muted/30 transition-colors group relative"
          >
            <!-- Main Row -->
            <div class="p-4 flex items-center gap-4 min-h-[4.5rem]">
              
              <!-- Label Section -->
              <div class="flex-1 min-w-0 flex items-center gap-3">
                 <div class="flex flex-col">
                    <label class="text-sm font-medium text-foreground truncate select-text">{{ line.key }}</label>
                    <span v-if="line.metadata?.type" class="text-[10px] uppercase font-bold tracking-wider text-muted-foreground">{{ line.metadata.type }}</span>
                 </div>
                 
                 <!-- Info Button -->
                 <button 
                   v-if="line.metadata?.description"
                   @click="toggleHelp(line.index)"
                   class="h-6 w-6 rounded-full flex items-center justify-center transition-colors focus:outline-none focus:ring-2 focus:ring-ring"
                   :class="activeHelp === line.index ? 'bg-primary/10 text-primary' : 'text-muted-foreground/50 hover:text-foreground hover:bg-muted'"
                   :title="i18n.t.description"
                 >
                   <Info :size="14" />
                 </button>
              </div>

              <!-- Input Section -->
              <div class="w-1/2 max-w-xs flex justify-end">
                 <!-- Boolean Toggle -->
                 <div v-if="getInputType(line.value, line.metadata) === 'bool'">
                    <button 
                      @click="onUpdateCFG(line, getDisplayValue(line.index, line.key, line.value) === 'true' ? 'false' : 'true')"
                      class="relative h-7 w-12 rounded-full transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-background focus:ring-primary"
                      :class="getDisplayValue(line.index, line.key, line.value) === 'true' ? 'bg-primary' : 'bg-input'"
                    >
                      <span 
                        class="absolute left-1 top-1 h-5 w-5 bg-white rounded-full transition-transform duration-200 shadow-sm flex items-center justify-center"
                        :class="getDisplayValue(line.index, line.key, line.value) === 'true' ? 'translate-x-5' : 'translate-x-0'"
                      >
                         <Check v-if="getDisplayValue(line.index, line.key, line.value) === 'true'" :size="10" class="text-primary stroke-[3]" />
                         <X v-else :size="10" class="text-muted-foreground stroke-[3]" />
                      </span>
                    </button>
                 </div>
                 
                 <!-- Text/Number Input -->
                 <div v-else class="w-full relative">
                   <input 
                     :type="getInputType(line.value, line.metadata) === 'number' ? 'number' : 'text'"
                     :value="getDisplayValue(line.index, line.key, line.value)"
                     @input="e => onUpdateCFG(line, (e.target as HTMLInputElement).value)"
                     class="flex h-9 w-full rounded-md border border-input bg-background px-3 py-1 text-sm shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 text-right font-mono"
                     :placeholder="line.metadata?.default"
                   />
                 </div>
              </div>
            </div>

            <!-- Expandable Help Section -->
            <div 
              v-if="activeHelp === line.index" 
              class="px-4 pb-4 pt-0 text-sm text-muted-foreground animate-accordion-down overflow-hidden border-t border-dashed border-border/50 bg-muted/20"
            >
              <div class="pt-3 flex gap-3">
                 <div class="mt-0.5 text-primary shrink-0">
                    <AlertCircle :size="16" />
                 </div>
                 <div class="flex-1">
                    <p class="leading-relaxed">{{ getDescription(line.key, line.metadata.description) }}</p>
                    <div class="mt-2 flex items-center gap-2 text-xs">
                       <span class="font-semibold text-foreground">Default:</span>
                       <code class="px-1 py-0.5 rounded bg-muted text-muted-foreground font-mono">{{ line.metadata.default }}</code>
                    </div>
                 </div>
              </div>
            </div>
          </div>
        </template>

        <!-- XML Content -->
      <div v-else-if="isXML" class="max-w-6xl mx-auto pb-20 px-4">
         
         <!-- Development Warning Banner -->
         <div class="mb-6 p-4 bg-red-500/10 border border-red-500/20 rounded-lg flex items-center gap-3">
            <div class="p-2 bg-red-500/20 rounded-full shrink-0">
               <AlertTriangle class="w-5 h-5 text-red-500" />
            </div>
            <span class="text-sm font-medium text-red-500 leading-tight">{{ i18n.t.xmlDevWarning }}</span>
         </div>

         <div class="space-y-2">
            <template v-for="line in xmlLines" :key="line.index">
              <!-- Only show lines that have attributes or look like structural tags -->
              <!-- Render logic: Indent based on depth. If attributes exist, show card. If not, maybe show simple tag? -->
              
              <!-- Structural Tag (No attributes) - Just for context -->
              <div v-if="!line.segments.some(s => s.isAttribute) && line.tagName" 
                   class="flex items-center text-slate-500 font-mono text-xs py-1 select-none"
                   :style="{ marginLeft: `${line.depth * 0.8}rem` }">
                  <span class="opacity-50">&lt;</span>
                  <span class="text-sky-700/70 font-bold">{{ line.tagName }}</span>
                  <span class="opacity-50">&gt;</span>
              </div>

              <!-- Editable Tag (Has attributes) -->
              <div v-else-if="line.segments.some(s => s.isAttribute)" 
                   class="relative group transition-all duration-300"
                   :style="{ marginLeft: `${line.depth * 0.8}rem` }">
                   
                  <!-- Connecting line to parent (visual sugar) -->
                  <div class="absolute -left-4 top-0 bottom-0 w-px bg-slate-800/50 group-hover:bg-slate-700/50 transition-colors"></div>

                  <div class="bg-slate-900/40 hover:bg-slate-900/80 border border-slate-800/60 hover:border-sky-500/30 rounded-lg p-3 shadow-sm hover:shadow-md transition-all">
                    
                    <!-- Header with Tag Name -->
                    <div class="flex items-center gap-3 mb-3 border-b border-slate-800/50 pb-2">
                        <span class="text-xs font-bold font-mono text-sky-400 bg-sky-950/30 px-2 py-0.5 rounded border border-sky-900/50">{{ line.tagName || 'Attributes' }}</span>
                        <span class="text-[10px] text-slate-600 font-mono">Line {{ line.index + 1 }}</span>
                    </div>

                    <!-- Attributes Grid -->
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                       <div v-for="seg in line.segments.filter(s => s.isAttribute)" :key="seg.attrName" class="group/field">
                          <label class="block text-[11px] font-bold text-slate-400 mb-1.5 uppercase tracking-wider group-hover/field:text-sky-300 transition-colors">{{ seg.attrName }}</label>
                          <div class="relative">
                            <input 
                                type="text" 
                                :value="getDisplayValue(line.index, seg.attrName!, seg.attrValue!)"
                                @input="e => onUpdateXML(line, seg.attrName!, (e.target as HTMLInputElement).value)"
                                class="w-full bg-slate-950 border border-slate-800 rounded px-3 py-2 text-slate-200 text-sm font-mono focus:outline-none focus:border-sky-500 focus:ring-1 focus:ring-sky-500/50 transition-all placeholder-slate-700"
                            />
                            <!-- Optional boolean toggle for "true"/"false" values -->
                            <button 
                                v-if="['true', 'false'].includes(getDisplayValue(line.index, seg.attrName!, seg.attrValue!).toLowerCase())"
                                @click="onUpdateXML(line, seg.attrName!, getDisplayValue(line.index, seg.attrName!, seg.attrValue!) === 'true' ? 'false' : 'true')"
                                class="absolute right-1 top-1 bottom-1 px-2 rounded hover:bg-slate-800 text-xs font-bold uppercase transition-colors"
                                :class="getDisplayValue(line.index, seg.attrName!, seg.attrValue!) === 'true' ? 'text-green-500' : 'text-red-500'"
                            >
                                {{ getDisplayValue(line.index, seg.attrName!, seg.attrValue!) }}
                            </button>
                          </div>
                       </div>
                    </div>
                  </div>
              </div>
            </template>
         </div>
      </div>
      </div>

    </div>
  </div>
</template>
