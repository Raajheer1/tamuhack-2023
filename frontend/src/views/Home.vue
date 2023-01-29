<template>
  <div class="flex flex-col h-screen justify-between">
    <div class="h-full pt-12">
      <CTA />
      <div class="px-8">
        <p class="text-3xl">Get out and about</p>
        <div class="flex" @click="router.push('search')">
          <p class="text-sm mt-1 mr-1 tracking-wide">Book a flight with TimelyTravel</p>
          <i class="fa-solid fa-arrow-right mt-1.5"></i>
        </div>

        <div @click="router.push('search')" class="flex bg-zinc-200 rounded-full py-3 px-3 mt-10 justify-between">
          <div class="flex">
            <i class="fa-solid fa-magnifying-glass mt-1"></i>
            <p class="pl-2">Where to?</p>
          </div>
          <div class="bg-zinc-700 text-white text-xs rounded-full px-2 py-1">
            Random
          </div>
        </div>

        <div class="mt-10">
          <h2 class="text-xl">
            <i class="fa-solid fa-plane-departure"></i>
            Upcoming flights
          </h2>
          <FlightInfo :flight="plane" />
        </div>

        <div class="mt-10" >
          <div class="flex">
            <h2 class="text-xl">
              <i class="fa-solid fa-plane-departure"></i>
              Popular Vacation Spots
            </h2>
          </div>
          <div class="mt-4 ml-4 underline decoration-yellow-500 decoration-2">
            <i class="fa-solid fa-trophy text-yellow-500"></i>
            <span class="text-lg mx-2">Anchorage, Alaska</span>
          </div>
          <div class="mt-2 ml-4 underline decoration-slate-400 decoration-2">
            <i class="fa-solid fa-trophy text-slate-400"></i>
            <span class="text-lg mx-2">The Valley, Anguilla</span>
          </div>
          <div class="mt-2 ml-4 underline decoration-amber-600 decoration-2">
            <i class="fa-solid fa-trophy text-amber-600"></i>
            <span class="text-lg mx-2">South Island, New Zealand</span>
          </div>
        </div>

      </div>
    </div>
    <Footer>Footer</Footer>
  </div>
</template>

<script setup lang="ts">
import {Airport, Flight} from "@/types";
import FlightInfo from "@/components/FlightInfo.vue";
import Footer from "@/views/partials/Footer.vue";
import CTA from "@/components/CTA.vue";
import {useRouter} from "vue-router";
import useUserStore from "@/store/user";
import {onMounted} from "vue";

const userStore = useUserStore();
const router = useRouter();


const dep: Airport = {
  code: "SEA",
  name: "Seattle-Tacoma International Airport",
  city: "Seattle"
}

const arr: Airport = {
  code: "LAX",
  name: "Los Angeles International Airport",
  city: "Los Angeles"
}

const plane: Flight = {
  flight_number: `AA976`,
  departure: dep,
  arrival: arr,
  date: `05 Feb`,
  departure_time: `14:20`,
  arrival_time: `16:07`,
  delay_percentage: 11,
  delay_time: 20,
  price: 245
};

onMounted(() => {
  if(userStore.hasFetched != true) {
    userStore.fetchUser();
  }
})
</script>

<style scoped>

</style>