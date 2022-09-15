<template>
    <div class="h-screen w-screen" v-if="!loading">
        <router-view></router-view>
    </div>
</template>
<script setup lang="ts">
import Auth from './services/auth';

const loading = Auth.loading()
</script>
