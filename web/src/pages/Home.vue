<template>
    <div class="flex h-full w-full" v-if="authenticated">
        <div class="m-auto bg-slate-800 p-4 rounded rounded-lg">
            <div>
                <button @click="logout">Logout</button>
            </div>
            <h1 class="text-lg">Discord</h1>
            <button @click="login">Login</button>
        </div>
    </div>
</template>
<script lang="ts" setup>
import { login } from '@/services/discord';
import Auth from '@/services/auth';

const logout = Auth.signOut;

const authenticated = Auth.isAuthenticated();
if (!authenticated.value) {
    Auth.signIn();
}
</script>