import { defineStore } from 'pinia';
import { ref, computed } from 'vue';

type Locale = 'pt-BR' | 'en-US' | 'es-ES';

const translations = {
  'pt-BR': {
    title: 'Terje Config',
    selectFile: 'Selecione um arquivo',
    saveChanges: 'Salvar Alterações',
    saving: 'Salvando...',
    line: 'Linha',
    type: 'Tipo',
    default: 'Padrão',
    description: 'Descrição',
    selectModule: 'Selecione um módulo para configurar',
    chooseFile: 'Escolha um arquivo no menu lateral para começar a editar.',
    search: 'Buscar...',
    xmlDevWarning: 'Aviso: A edição de XML está em desenvolvimento. Em breve haverá um editor personalizado para cada arquivo.',
  },
  'en-US': {
    title: 'Terje Config',
    selectFile: 'Select a file',
    saveChanges: 'Save Changes',
    saving: 'Saving...',
    line: 'Line',
    type: 'Type',
    default: 'Default',
    description: 'Description',
    selectModule: 'Select a module to configure',
    chooseFile: 'Choose a file from the sidebar menu to start editing.',
    search: 'Search...',
    xmlDevWarning: 'Warning: XML editing is under development. A custom editor for each file will be available soon.',
  },
  'es-ES': {
    title: 'Terje Config',
    selectFile: 'Seleccione un archivo',
    saveChanges: 'Guardar Cambios',
    saving: 'Guardando...',
    line: 'Línea',
    type: 'Tipo',
    default: 'Defecto',
    description: 'Descripción',
    selectModule: 'Seleccione un módulo para configurar',
    chooseFile: 'Elija un archivo del menú lateral para comenzar a editar.',
    search: 'Buscar...',
    xmlDevWarning: 'Aviso: La edición de XML está en desarrollo. Pronto habrá un editor personalizado para cada archivo.',
  }
};

export const useI18nStore = defineStore('i18n', () => {
  const currentLocale = ref<Locale>('pt-BR');

  const t = computed(() => translations[currentLocale.value]);

  function setLocale(locale: Locale) {
    currentLocale.value = locale;
  }

  return {
    currentLocale,
    t,
    setLocale
  };
});
