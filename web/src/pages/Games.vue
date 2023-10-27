<script lang="ts" setup>
import { nextTick, onMounted, ref } from "vue";
import { Converter } from "showdown";

const html = ref<string>();
const games = ref<HTMLDivElement>();

onMounted(async () => {
    const res = await fetch("/games.md");
    const md = await res.text();

    const converter = new Converter({
        tables: true,
        tablesHeaderId: true,
    });
    html.value = converter.makeHtml(md);

    await nextTick();

    games.value?.scrollIntoView({ behavior: "smooth", block: "start" });
});
</script>

<template>
    <div
        ref="games"
        class="rounded pt-4 prose prose-md prose-invert text-left prose-neutral"
        v-html="html"
    ></div>
</template>
