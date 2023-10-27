<script lang="ts" setup>
import { nextTick, onMounted, ref } from "vue";
import { Converter } from "showdown";

const html = ref<string>();
const element = ref<HTMLDivElement>();

onMounted(async () => {
    const res = await fetch("/rules.md");
    const md = await res.text();

    const converter = new Converter({
        tables: true,
        tablesHeaderId: true,
    });
    html.value = converter.makeHtml(md);

    await nextTick();

    element.value?.scrollIntoView({ behavior: "smooth", block: "start" });
});
</script>

<template>
    <div
        ref="element"
        class="rounded pt-4 prose prose-md prose-invert text-left prose-neutral"
        v-html="html"
    ></div>
</template>
