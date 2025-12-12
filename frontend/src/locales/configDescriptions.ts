// This file contains translations for the configuration descriptions
// Keys must match exactly the key in the .cfg file

export const configDescriptions: Record<string, { 'pt-BR': string; 'en-US': string; 'es-ES': string; 'ru-RU': string }> = {
  // Core.cfg
  'Core.DatabaseAutosaveInterval': {
    'en-US': 'Once in a specified time, checks all player profiles for changes and, if necessary, saves them to disk. (in seconds)',
    'pt-BR': 'Uma vez em um tempo especificado, verifica todos os perfis de jogadores por alterações e, se necessário, salva-os no disco. (em segundos)',
    'es-ES': 'Una vez en un tiempo especificado, verifica todos los perfiles de jugadores para cambios y, si es necesario, los guarda en el disco. (en segundos)',
    'ru-RU': 'Раз в указанное время проверяет все профили игроков на наличие изменений и, при необходимости, сохраняет их на диск. (в секундах)'
  },
  'Core.ProfileSynchInterval': {
    'en-US': 'Once in a specified time, checks all player profiles for changes and, if necessary, send them to client side. (in seconds)',
    'pt-BR': 'Uma vez em um tempo especificado, verifica todos os perfis de jogadores por alterações e, se necessário, envia-os para o lado do cliente. (em segundos)',
    'es-ES': 'Una vez en un tiempo especificado, verifica todos los perfiles de jugadores para cambios y, si es necesario, los envía al lado del cliente. (en segundos)',
    'ru-RU': 'Раз в указанное время проверяет все профили игроков на наличие изменений и, при необходимости, отправляет их на клиентскую сторону. (в секундах)'
  },
  'Core.StatsSynchInterval': {
    'en-US': 'Once in a specified time, checks all player stats for changes and, if necessary, sync them with client side. (in seconds)',
    'pt-BR': 'Uma vez em um tempo especificado, verifica todas as estatísticas dos jogadores por alterações e, se necessário, sincroniza com o lado do cliente. (em segundos)',
    'es-ES': 'Una vez en un tiempo especificado, verifica todas las estadísticas de los jugadores para cambios y, si es necesario, las sincroniza con el lado del cliente. (en segundos)',
    'ru-RU': 'Раз в указанное время проверяет всю статистику игроков на наличие изменений и, при необходимости, синхронизирует ее с клиентской стороной. (в секундах)'
  },
  'Core.SatBinaryTreeGridSize': {
    'en-US': 'Sets the grid size of the binary search tree binary grid for intersecting scriptable areas. Do not change this value unnecessarily.',
    'pt-BR': 'Define o tamanho da grade da árvore de busca binária para áreas programáveis de interseção. Não altere este valor desnecessariamente.',
    'es-ES': 'Establece el tamaño de la cuadrícula del árbol de búsqueda binaria para áreas programables de intersección. No cambie este valor innecesariamente.',
    'ru-RU': 'Устанавливает размер сетки двоичного дерева поиска для пересекающихся скриптовых областей. Не изменяйте это значение без необходимости.'
  },
  'Core.WaterDrainFromVomit': {
    'en-US': 'The value of water units that a player will lose when vomiting.',
    'pt-BR': 'O valor de unidades de água que um jogador perderá ao vomitar.',
    'es-ES': 'El valor de unidades de agua que un jugador perderá al vomitar.',
    'ru-RU': 'Количество единиц воды, которое игрок потеряет при рвоте.'
  },
  'Core.EnergyDrainFromVomit': {
    'en-US': 'The value of energy units that a player will lose when vomiting.',
    'pt-BR': 'O valor de unidades de energia que um jogador perderá ao vomitar.',
    'es-ES': 'El valor de unidades de energía que un jugador perderá al vomitar.',
    'ru-RU': 'Количество единиц энергии, которое игрок потеряет при рвоте.'
  },
  'Core.FixKillOnDisconnect': {
    'en-US': 'Kills a player when he disconnects from the server while unconscious or restrained. Enabling this setting fixes abuse of mods logic.',
    'pt-BR': 'Mata o jogador quando ele se desconecta do servidor enquanto inconsciente ou algemado. Habilitar esta configuração corrige abusos de lógica de mods.',
    'es-ES': 'Mata a un jugador cuando se desconecta del servidor mientras está inconsciente o restringido. Habilitar esta configuración corrige el abuso de la lógica de mods.',
    'ru-RU': 'Убивает игрока, если он отключается от сервера, находясь без сознания или будучи связанным. Включение этой настройки исправляет злоупотребление логикой модов.'
  },
  
  // Medicine.cfg
  'Medicine.MindCanSuicide': {
    'en-US': 'This parameter determines whether the player can try to suicide when their sanity (mind) level is critical.',
    'pt-BR': 'Este parâmetro determina se o jogador pode tentar suicídio quando seu nível de sanidade (mente) for crítico.',
    'es-ES': 'Este parámetro determina si el jugador puede intentar suicidarse cuando su nivel de cordura (mente) es crítico.',
    'ru-RU': 'Этот параметр определяет, может ли игрок попытаться совершить самоубийство, когда уровень его рассудка (разума) критический.'
  },
  'Medicine.BotKillingMindDegradationValue': {
    'en-US': 'The amount of mind that a player will lose per second after killing an AI (bot).',
    'pt-BR': 'A quantidade de mente que um jogador perderá por segundo após matar uma IA (bot).',
    'es-ES': 'La cantidad de mente que un jugador perderá por segundo después de matar una IA (bot).',
    'ru-RU': 'Количество рассудка, которое игрок будет терять в секунду после убийства ИИ (бота).'
  },
  
  // Skills.cfg
  'Skills.SurvLifetimeOffset': {
    'en-US': 'Sets the cyclic time interval in seconds after which the player will be awarded additional experience points for the survival skill.',
    'pt-BR': 'Define o intervalo de tempo cíclico em segundos após o qual o jogador receberá pontos de experiência adicionais para a habilidade de sobrevivência.',
    'es-ES': 'Establece el intervalo de tiempo cíclico en segundos después del cual el jugador recibirá puntos de experiencia adicionales para la habilidad de supervivencia.',
    'ru-RU': 'Устанавливает циклический интервал времени в секундах, по истечении которого игроку будут начислены дополнительные очки опыта за навык выживания.'
  }
};
