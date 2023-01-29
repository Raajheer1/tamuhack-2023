<template>
  <SearchAndBookFlight v-if="search" />
  <Billing v-if="search" />
</template>

<script setup lang="ts">
import SearchAndBookFlight from "@/components/SearchAndBookFlight.vue";
import Billing from "@/components/Billing.vue";
import {ref} from "vue";

import useSearchStore from "@/store/search";
import SearchResults from "@/views/partials/SearchResults.vue";
const SearchStore = useSearchStore();

const search = ref(false);
</script>

<style scoped>

</style>