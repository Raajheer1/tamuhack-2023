<template>
  <SearchAndBookFlight v-if="search" />
</template>

<script setup lang="ts">
import SearchAndBookFlight from "@/components/SearchAndBookFlight.vue";
import {ref} from "vue";

const search = ref(false);
</script>

<style scoped>

</style>