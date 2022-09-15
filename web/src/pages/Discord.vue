<template>
    <div></div>
</template>
<script lang="ts" setup>
import { useRoute } from 'vue-router';

const route = useRoute();

const code = route.query.code as string | null

if (code) {

}
</script>