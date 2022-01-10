
<script setup>
import { KeepAlive, Transition } from 'vue'

</script>


<template>
    <div>
        <router-view v-slot="{ Component }">
            <transition mode="out-in" name="el-fade-in-linear">
                <keep-alive>
                    <component :is="Component" />
                </keep-alive>
            </transition>
        </router-view>
    </div>
</template>