<template>
    <div class="h-screen w-screen" v-if="!isLoading">
        <router-view></router-view>
    </div>
</template>
<script setup lang="ts">
import { useAuth0 } from "@auth0/auth0-vue";

const { isLoading } = useAuth0();
</script>
