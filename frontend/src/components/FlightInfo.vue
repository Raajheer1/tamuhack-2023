<template>
  <div class="bg-slate-100 pl-2 pt-1.5 rounded">
    <div class="grid grid-cols-10 gap-x-5">
      <div class="flex col-span-10">
        <div class="grid grid-rows-2 max-w-[30px]">
          <span class="text-zinc-500 text-xl mx-2">{{ flight.date.substring(0,2) }}</span>
          <span class="text-zinc-500 text-sm ml-1.5 -mt-1 text-clip">{{ flight.date.substring(2) }}</span>
        </div>
        <img src="@/assets/aa.png" alt="American Airlines" class="h-10 pt-2 pr-1" />
        <div class="grid grid-cols-10">
          <div class="grid grid-rows-2">
            <p class="text-md">{{ flight.departure_time }}</p>
            <p class="text-sm text-zinc-500 -mt-1.5">{{ flight.departure.code }}</p>
          </div>
          <div class="col-span-7 mx-3 text-center">
            <div class="grid grid-rows-2 h-5/6">
              <div class="border-dotted border-b-2 border-black">
                <span class="text-emerald-500 text-xs mr-2">NON-STOP</span>
                <span class="text-zinc-500 text-xs ml-2">2H 42M</span>
              </div>
              <div>
                <p v-if="flight.delay_percentage < 5" class="text-xs text-zinc-900 underline decoration-emerald-500">
                  <i class="fa-solid fa-check text-emerald-500"></i>
                  Often on time
                </p>
                <p v-else-if="flight.delay_percentage < 15" class="text-xs text-zinc-900 underline decoration-yellow-500">
                  <i class="fa-solid fa-triangle-exclamation text-yellow-500"></i>
                  Possible Delay
                </p>
                <p v-else class="text-xs text-zinc-900 underline decoration-red-500">
                  <i class="fa-solid fa-circle-exclamation text-red-500"></i>
                  Often Delayed
                </p>
              </div>
            </div>

          </div>
          <div class="grid grid-rows-2 text-right">
            <p class="text-md">{{ flight.arrival_time }}</p>
            <p class="text-sm text-zinc-500 -mt-1.5">{{ flight.arrival.code }}</p>
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