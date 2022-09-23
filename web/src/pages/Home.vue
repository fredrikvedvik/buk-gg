<template>
    <div class="flex h-full w-full" v-if="authenticated">
        <div class="fixed top-0 right-0" v-if="user?.isAdmin">
            <Admin></Admin>
        </div>
        <div class="m-auto bg-slate-800 p-4 rounded rounded-lg">
            <div>
                <button @click="logout">Logout</button>
            </div>
            <Discord v-if="settings !== null" :settings="settings"></Discord>
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

const logout = Auth.signOut;

const authenticated = Auth.isAuthenticated();
if (!authenticated.value) {
    Auth.signIn();
} else {
    api.getUser().then(r => {
        user.value = r
    })
    api.getSettings().then(r => {
        settings.value = r
    });
}
</script>