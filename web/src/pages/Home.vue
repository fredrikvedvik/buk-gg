<template>
    <div class="flex h-full w-full" v-if="authenticated">
        <div class="fixed top-0 right-0" v-if="user?.isAdmin">
            <Admin></Admin>
        </div>
        <div class="m-auto">
            <img class="invert mb-8 mx-auto" src="/logo.svg" />
            <div class="bg-slate-800 p-8 rounded rounded-lg">
                <div class="flex mb-2 gap-4">
                    <p class="my-auto text-md">Logged in as <span class="text-lg">{{u.name}}</span></p>
                    <button class="ml-auto bg-red-500 p-2 rounded px-3" @click="logout">Logout</button>
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