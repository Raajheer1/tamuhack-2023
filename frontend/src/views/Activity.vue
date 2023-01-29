<template>
  <div class="flex flex-col h-screen justify-between">
    <div class="h-full pt-8 mt-4">
      <div class="px-8 mb-5">
        <h1 class="text-4xl">Activity Log</h1>
        <div class="flex justify-between">
          <h2 class="text-2xl">Past</h2>
          <span @click="filtersMenu = !filtersMenu" class="flex bg-zinc-200 p-1 px-3 rounded-full cursor-pointer">
            <i class="fa-solid fa-sliders p-1"></i>
            <p class="pl-0.5">Filters</p>
          </span>
        </div>
        <div class="flex" v-if="filtersMenu == true">
          Some Filter options here...
          Maybe by airline or by date?
        </div>
      </div>

      <div class="flex border-t-2 border-zinc-20" v-for="x in 5">
        <img src="@/assets/aa.png" class="w-20" />
        <div class="w-full grid grid-cols-4">
          <div class="col-span-3">
            <p>AAL174 <span class="text-zinc-500">ORD - LAX</span></p>
            <p class="text-xs text-zinc-500 -mt-0.5">Jun 21st</p>
            <p class="text-xs text-zinc-500">$1,500</p>
          </div>
          <div class="mt-4 mr-3 text-right">
            <i class="fa-solid fa-arrow-right"></i>
          </div>
        </div>
      </div>
    </div>
    <Footer>Footer</Footer>
  </div>
</template>

<script setup lang="ts">
import FlightInfo from "@/components/FlightInfo.vue";
import Footer from "@/views/partials/Footer.vue";
import CTA from "@/components/CTA.vue";
import {useRouter} from "vue-router";
import useUserStore from "@/store/user";
import {ref} from "vue";

const userStore = useUserStore();
const router = useRouter();

const filtersMenu = ref(false);

</script>

<style scoped>

</style>