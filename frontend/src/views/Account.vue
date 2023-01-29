<template>
  <div class="flex flex-col h-screen justify-between">
    <div class="h-full">
      <div class="px-8 pt-12">
        <div class="flex justify-between">
          <h1 class="text-3xl mt-1">Raaj Patel</h1>
          <i class="fa-solid fa-user rounded-full border border-4 border-black p-3"></i>
        </div>
        <div class="grid grid-cols-3 py-5 gap-x-5 border-b-2 border-black">
          <div class="flex items-center text-center bg-zinc-200 rounded-lg">
            <div class="col-span-1 w-full py-4">
              <div @click="switchModal('help')" class="grid grid-rows-2">
                <i class="fa-solid fa-headset"></i>
                <p class="text-sm">Help</p>
              </div>
            </div>
          </div>
          <div class="flex items-center text-center bg-zinc-200 rounded-lg">
            <div class="col-span-1 w-full py-4">
              <div @click="switchModal('wallet')" class="grid grid-rows-2">
                <i class="fa-solid fa-wallet"></i>
                <p class="text-sm">Wallet</p>
              </div>
            </div>
          </div>
          <div class="flex items-center text-center bg-zinc-200 rounded-lg">
            <div class="col-span-1 w-full py-4">
              <div @click="switchModal('trips')" class="grid grid-rows-2">
                <i class="fa-solid fa-suitcase"></i>
                <p class="text-sm">Trips</p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <Wallet v-if="info === 'wallet'" />
      <Help v-if="info === 'help'" />
      <Trips v-if="info === 'trips'" />
    </div>
    <Footer>Footer</Footer>
  </div>
</template>

<script setup lang="ts">
import FlightInfo from "@/components/FlightInfo.vue";
import Help from "@/components/Help.vue";
import Wallet from "@/components/Wallet.vue";
import Trips from "@/components/Trips.vue"
import Footer from "@/views/partials/Footer.vue";
import CTA from "@/components/CTA.vue";
import {useRouter} from "vue-router";
import useUserStore from "@/store/user";
import { ref } from "vue";

const userStore = useUserStore();
const router = useRouter();

const info = ref('trips')

function switchModal(modal:string) {
  info.value = modal
}

</script>

<style scoped>
</style>