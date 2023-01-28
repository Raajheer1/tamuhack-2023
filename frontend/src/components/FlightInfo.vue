<template>
  <div class="bg-slate-100 pl-2 py-3 rounded">
    <div class="grid grid-cols-10 gap-x-5">
      <div class="flex col-span-6">
        <img src="@/assets/aa.png" alt="American Airlines" class="w-14 pr-1" />
        <div class="grid grid-rows-2">
          <div class="flex">
            <p class="text-sm">{{ flight.departure.city }}</p>
            <i class="fa-solid fa-arrow-right mx-2 mt-0.5 text-red-500"></i>
            <p class="text-sm">{{ flight.arrival.city }}</p>
          </div>
          <div class="flex py-1">
            <div class="bg-slate-200 rounded-full">
              <p class="text-xs px-2 text-zinc-700">{{ flight.departure_time }}</p>
            </div>
          </div>
        </div>
      </div>
      <div class="col-span-4">
        <div class="grid grid-rows-2">
          <p class="text-sm">{{ flight.date }}</p>
          <div class="flex py-1">
            <p class="text-xs text-zinc-900 underline decoration-red-500">
              <i class="fa-solid fa-triangle-exclamation text-red-500"></i>
              Often Delayed
            </p>
          </div>
        </div>
      </div>

    </div>

  </div>
</template>

<script setup lang="ts">
import {Flight} from "@/types";

defineProps<{ flight: Flight }>()

</script>

<style scoped>

</style>