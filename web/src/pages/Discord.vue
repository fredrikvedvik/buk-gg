<template>
    <div>
        <h1 class="text-lg">Discord</h1>
        <div v-for="discord in settings.discords">
            <h1>{{discord.username}}#{{discord.discriminator}} <span class="cursor-pointer" @click="removeDiscord(discord.id)">X</span></h1>
        </div>
        <button @click="login" v-if="!settings.discords || settings.discords.length <= 1">Login</button>
    </div>
</template>
<script lang="ts" setup>
import { Settings } from '@/services/api/types';
import { login, removeDiscord } from '@/services/discord';

defineProps<{
    settings: Settings
}>()
</script>