<template>
    <div class="flex flex-col bg-slate-700 p-4 rounded rounded-lg gap-4">
        <h1 class="text-lg">Linked accounts</h1>
        <div class="flex flex-col gap-2">
            <div v-for="discord in settings.discords" class="flex bg-slate-800 p-2 rounded rounded-lg px-4">
                <h1 class="my-auto">{{discord.username}}#{{discord.discriminator}}</h1>
                <XMarkIcon class="h-8 w-8 ml-auto fill-red-300 hover:-translate-y-0.5 transition duration-200 cursor-pointer" @click="removeDiscord(discord.id)"></XMarkIcon>
            </div>
        </div>
        <button @click="login" class="mt-2 bg-green-500 p-2 rounded rounded-lg px-3" v-if="!settings.discords || settings.discords.length <= 1">Add</button>
        <p class="text-gray-500" v-if="settings.discords?.length >= 2">You can only have 2 accounts connected to your account</p>
    </div>
</template>
<script lang="ts" setup>
import api from '@/services/api';
import { Settings } from '@/services/api/types';
import { login, removeDiscord } from '@/services/discord';
import { XMarkIcon } from "@heroicons/vue/24/solid"

defineProps<{
    settings: Settings
}>()

// api.getGuilds().then(r => {
//     console.log(r)
// })
</script>