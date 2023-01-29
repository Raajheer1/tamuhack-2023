<template>
  <div class="flex flex-col justify-between h-screen bg-slate-100 pl-2 py-3">
    <div class="flex">
      <img src="@/assets/aa.png" alt="American Airlines" class="h-20 pt-2 pr-1" />
      <div class="grid grid-rows-2 mt-3">
        <h1 class="text-2xl">American Airlines - AA</h1>
        <div class="flex gap-x-4">
          <h2 class="text-xl"><i class="fa-solid fa-plane"></i> {{ flight.flight_number }}</h2>
          <h2 class="text-xl"><i class="fa-solid fa-calendar"></i> {{ flight.date }}</h2>
        </div>
      </div>
    </div>
    <div class="h-full flex flex-col justify-between">
      <div class="text-md text-zinc-500">
        <span class="text-sm text-black"><i class="fa-solid fa-plane-departure"></i> Departure</span>
        <div class="flex text-black text-2xl">
          {{ flight.departure.code }} -
          {{ flight.departure.city}}
        </div>
        {{ flight.departure.name }}
      </div>

      <div class="grid grid-cols-8 h-full">
        <div class="border-dotted border-r-2 border-black">
        </div>
        <div class="m-4 col-span-7">
          <h3 class="text-lg underline">Delay Information:</h3>
          <div class="flex gap-x-2 mt-2">
            <i v-if="flight.delay_percentage < 5" class="fa-solid fa-check text-emerald-500"></i>
            <i v-else-if="flight.delay_percentage < 15" class="fa-solid fa-triangle-exclamation text-yellow-500"></i>
            <i v-else class="fa-solid fa-circle-exclamation text-red-500"></i>
            <p class="text-sm -m-0.5">{{ flight.delay_percentage }}% of flights delayed</p>
          </div>
          <div class="flex gap-x-3 mt-1">
            <i v-if="flight.delay_time < 15" class="fa-solid fa-hourglass text-emerald-500"></i>
            <i v-else-if="flight.delay_time < 30" class="fa-solid fa-hourglass text-yellow-500"></i>
            <i v-else class="fa-solid fa-hourglass text-red-500"></i>
            <p class="text-sm -m-0.5">{{ flight.delay_time }} minutes average delay</p>
          </div>
          <h3 class="text-lg mt-4 underline">Flight Information:</h3>
          <p class="text-sm">Non-Stop</p>
          <p class="text-sm">Aircraft: B73M</p>
          <p class="text-sm">ETA: {{ flight.arrival_time }}</p>
          <p class="text-sm">Gate: Terminal 2, Gate 41</p>
        </div>
      </div>

      <div class="text-md text-zinc-500">
        <span class="text-sm text-black"><i class="fa-solid fa-plane-arrival"></i> Arrival</span>
        <div class="flex text-black text-2xl">
          {{ flight.arrival.code }} -
          {{ flight.arrival.city}}
        </div>
        {{ flight.arrival.name }}
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