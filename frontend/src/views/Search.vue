<template>
  <SearchAndBookFlight v-if="search" />
  <Billing v-if="search" />
</template>

<script setup lang="ts">
import SearchAndBookFlight from "@/components/SearchAndBookFlight.vue";
import Billing from "@/components/Billing.vue";
import {ref} from "vue";

const search = ref(true);
</script>

<style scoped>

</style>