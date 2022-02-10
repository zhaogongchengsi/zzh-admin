<script setup lang="ts">
import { 
  IconMinus, 
  IconClose, 
  IconFullscreen, 
  IconFullscreenExit,  
  IconMenuFold,
  IconMenuUnfold,
  IconCommand,
  IconSunFill,
  IconMoonFill
} from '@arco-design/web-vue/es/icon';
import { ref } from 'vue'
import { useMenuStore } from '@/pinia'
import sendMessage from '@/sysetm';
const menuStore = useMenuStore()
// const { sendMessage } = window.blogApp;
enum sysetm {
    mini = "onMinimize",
    full = "onFullScreen",
    fuit = "onFullScreenExit",
    cose = "onClose",
}
const isFullscreen = ref(true)
const onSettings = (type:sysetm) => {
  if (type === 'onFullScreen') {
    isFullscreen.value = false;
  } else if (type === 'onFullScreenExit') {
    isFullscreen.value = true;
  }
  sendMessage(type)
}

const handle = () =>{
}

</script>
<template>
  <div :class="['toptool-container', {'toptool-container-dark': menuStore.theme}]">
    <div class="toptool-left"
      @click="menuStore.toggleCollapse"
    >
      <icon-menu-unfold v-if="menuStore.collapsed" />
      <icon-menu-fold v-else />
    </div>
    <div class="toptool-right">
      <blog-space size="large">
      <a-dropdown @select="handle">
        <icon-command />
        <template #content>
          <a-doption>
              <a-switch 
                type="large"
                checked-color="#333" 
                unchecked-color="#ccc"
                checked-value="dark" 
                unchecked-value="light"
                v-model="menuStore.theme"
                @change="menuStore.setTheme"
              >
                <template #checked>  
                  <a-typography-text>
                    dark
                  </a-typography-text>
                </template>
                <template #unchecked>
                  <a-typography-text>
                    light
                  </a-typography-text>
                </template>
                <template #checked-icon>
                  <icon-moon-fill />
                </template>
                <template #unchecked-icon>
                  <icon-sun-fill />
                </template>
              </a-switch>
          </a-doption>
          <a-doption disabled>Option 2</a-doption>
          <a-doption :value="{ value: 'Option3' }">Option 3</a-doption>
        </template>
      </a-dropdown>
      <icon-minus @click="onSettings(sysetm.mini)" />
      <icon-fullscreen v-if="isFullscreen" @click="onSettings(sysetm.full)" />
      <icon-fullscreen-exit v-else @click="onSettings(sysetm.fuit)" />
      <icon-close @click="onSettings(sysetm.cose)" />
    </blog-space>
    </div>
  </div>
</template>

<style scoped lang="scss">
.toptool-container {
  --tools-padding: 15px;
  height: var(--header-tools-height);
  -webkit-app-region: drag;
  display: flex;
  justify-content:space-between;
  border-bottom: 1px solid var(--color-neutral-3);
  .toptool-right {
    margin-right: var(--tools-padding);
    display: flex;
    align-items: center;
    justify-content: center;
  }
}

.toptool-left {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-left: var(--tools-padding);
  -webkit-app-region: no-drag;
}

.toptool-container:deep(.arco-icon) {
  -webkit-app-region: no-drag;
  cursor: pointer;
  width: 1.2rem;
  height: 1.2rem;
  border-radius: 3px;
  color: var(--color-text-3);
  &:hover {
    background-color: var(--color-secondary-hover);
    color: var(--color-text-2);
  }
}

.toptool-container-dark:deep(.arco-icon) {
  color: var(--color-text-1);;
}

</style>
