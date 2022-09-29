<template>
    <div class="flex flex-col bg-slate-700 p-4 rounded rounded-lg gap-2 shadow-xl">
        <h1 class="text-lg uppercase">Linked accounts</h1>
        <p class="text-gray-500">These are your linked <a href="https://discord.com" class="text-blue-500">Discord</a> accounts</p>
        <div class="flex flex-col gap-2">
            <div v-for="discord in settings.discords" class="flex bg-slate-800 p-2 rounded rounded-lg px-4">
                <div class="my-auto">
                    <h1 class="text-lg">{{discord.username}}#{{discord.discriminator}}</h1>
                    <p class="text-gray-500 text-sm">{{discord.id}}</p>
                </div>
                <XMarkIcon class="h-10 w-10 ml-auto my-auto fill-red-500 hover:fill-red-600 hover:-translate-y-0.5 transition duration-100 cursor-pointer" @click="removeDiscord(discord.id)"></XMarkIcon>
            </div>
        </div>
        <button @click="login" class="mt-2 bg-green-500 p-2 rounded rounded-lg px-3 hover:bg-green-600 transition duration-100 hover:-translate-y-0.5 uppercase" v-if="!settings.discords || settings.discords.length <= 1">Add account</button>
        <p class="text-gray-500" v-if="settings.discords?.length >= 2">You can only have 2 accounts connected to your account</p>
    </div>
    <div class="px-4 mt-2">
        <p class="text-gray-500">
            You should automatically be added to the server when connecting your account
        </p>
    </div>
</template>
<script lang="ts" setup>
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