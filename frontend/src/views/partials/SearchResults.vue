<template>
  <div class="text-center">
    <h1 class="text-xl">{{ origin.city }} ({{origin.code}}) to {{ destination.city }} ({{ destination.code }})</h1>
  </div>
  <div class="grid grid-cols-1 gap-y-1 mx-2" v-if="false">
    <FlightInfo v-for="flight in flights" :key="flight.flight_number" :flight="flight" />
  </div>
</template>

<script setup lang="ts">
import FlightInfo from "@/components/FlightInfo.vue";
import {Airport, Flight} from "@/types/index.js";

defineProps<{ flights: Flight[], origin: Airport, destination: Airport }>()


</script>

<style scoped>

</style>