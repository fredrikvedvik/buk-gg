<template>
    <div class="bg-slate-700 p-4 rounded rounded-lg">
        <h1 class="text-lg">Linked accounts</h1>
        <div v-for="discord in settings.discords" class="flex bg-slate-800 p-2 rounded rounded-lg px-4">
            <h1 class="my-auto">{{discord.username}}#{{discord.discriminator}}</h1>
            <XMarkIcon class="h-8 w-8 ml-auto fill-red-300 hover:-translate-y-0.5 transition duration-200 cursor-pointer" @click="removeDiscord(discord.id)"></XMarkIcon>
        </div>
        <button @click="login" class="mt-2 bg-green-500 p-2 rounded rounded-lg px-3" v-if="!settings.discords || settings.discords.length <= 1">Add</button>
    </div>
</template>
<script lang="ts" setup>
import { Settings } from '@/services/api/types';
import { login, removeDiscord } from '@/services/discord';
import { XMarkIcon } from "@heroicons/vue/24/solid"

defineProps<{
    settings: Settings
}>()
</script>