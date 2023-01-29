<template>
  <div class="flex flex-col h-screen justify-between">
    <i @click="router.push('/')" class="ml-4 mt-12 fa-solid fa-arrow-left"></i>
    <div class="h-full pt-8">
      <div class="px-8">
        <h1 class="text-4xl mb-2">Where to?</h1>
        <SearchAndBookFlight v-if="!search" />
      </div>
    </div>
  </div>
  <Billing v-if="search" />
</template>

<script setup lang="ts">
import SearchAndBookFlight from "@/components/SearchAndBookFlight.vue";
import Billing from "@/components/Billing.vue";
import {ref} from "vue";

import useSearchStore from "@/store/search";
import SearchResults from "@/views/partials/SearchResults.vue";
import {useRouter} from "vue-router";
const SearchStore = useSearchStore();

const router = useRouter();

const search = ref(false);
</script>

<style scoped>

</style>