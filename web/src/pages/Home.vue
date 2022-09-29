<template>
    <div class="flex h-full w-full" v-if="authenticated">
        <div class="fixed w-full h-full bg-cover bg-no-repeat z-0 blur-xl opacity-10" style="background-image: url('bg.svg')"></div>
        <div class="fixed top-0 right-0 z-20" v-if="user?.isAdmin">
            <Admin></Admin>
        </div>
        <div v-if="settings === null" wire:loading class="fixed top-0 left-0 right-0 bottom-0 w-full h-screen z-50 overflow-hidden opacity-75 flex flex-col items-center justify-center">
            <svg aria-hidden="true" class="w-16 mb-4 text-gray-200 animate-spin dark:text-gray-600 fill-blue-600" viewBox="0 0 100 101" fill="none" xmlns="http://www.w3.org/2000/svg">
                <path d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z" fill="currentColor"/>
                <path d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z" fill="currentFill"/>
            </svg>
            <h2 class="text-center text-white text-xl font-semibold">Loading...</h2>
            <p class="w-1/3 text-center text-white">This may take a few moments</p>
        </div>
        <div v-else class="m-auto z-10">
            <div class="bg-slate-800 p-8 rounded rounded-lg w-screen max-w-sm shadow-xl">
                <img class="invert mb-8 mx-auto" src="/logo.svg" />
                <div class="flex mb-2 gap-4">
                    <p class="my-auto text-md"><span class="text-lg">{{u.name}}</span></p>
                    <button class="ml-auto bg-red-500 p-2 rounded px-3 hover:bg-red-600 transition duration-100 hover:-translate-y-0.5 uppercase" @click="logout">Logout</button>
                </div>
                <Discord v-if="settings !== null" :settings="settings"></Discord>
                <div v-if="loading">Loading...</div>
            </div>
        </div>
    </div>
</template>
<script lang="ts" setup>
import api from '@/services/api';
import Auth from '@/services/auth';
import Discord from './Discord.vue';
import settings from "@/services/settings"
import user from "@/services/user"
import Admin from "./Admin.vue"
import { ref } from 'vue';

const logout = Auth.signOut;

const u = Auth.user();

const loading = ref(true)

const authenticated = Auth.isAuthenticated();
if (!authenticated.value) {
    Auth.signIn();
} else {
    api.getUser().then(r => {
        user.value = r

        loading.value = false
    })
    api.getSettings().then(r => {
        settings.value = r

        loading.value = false
    });
}
</script>