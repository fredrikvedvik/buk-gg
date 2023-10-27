<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { Converter } from "showdown";

const html = ref<string>();

onMounted(async () => {
    const res = await fetch("/rules.md");
    const md = await res.text();

    const converter = new Converter({
        tables: true,
        tablesHeaderId: true,
    });
    html.value = converter.makeHtml(md);
});
</script>

<template>
    <div
        class="rounded prose prose-md prose-invert text-left prose-neutral"
        v-html="html"
    ></div>
</template>
