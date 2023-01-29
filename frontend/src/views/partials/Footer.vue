<template>
  <div class="grid grid-cols-3 py-5 border-t-2">
    <div class="flex items-center text-center">
      <div class="col-span-1 w-full">
        <div @click="router.push('/')" class="grid grid-rows-2" :class="{'text-zinc-400':currentRoute!='Home'}">
          <i class="fa-solid fa-house"></i>
          <p class="text-sm">Home</p>
        </div>
      </div>
    </div>
    <div class="flex items-center text-center">
      <div class="col-span-1 w-full">
        <div @click="router.push('/activity')" class="grid grid-rows-2" :class="{'text-zinc-400':currentRoute!='Activity'}">
          <i class="fa-solid fa-receipt"></i>
          <p class="text-sm">Activity</p>
        </div>
      </div>
    </div>
    <div class="flex items-center text-center">
      <div class="col-span-1 w-full">
        <div @click="router.push('/account')" class="grid grid-rows-2" :class="{'text-zinc-400':currentRoute!='Account'}">
          <i class="fa-solid fa-user"></i>
          <p class="text-sm">Account</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {useRouter,useRoute} from "vue-router";
import {computed} from "vue";

const route = useRoute();
const router = useRouter()


const currentRoute = computed(() => {
  return route.name;
})
</script>

<style scoped>

</style>