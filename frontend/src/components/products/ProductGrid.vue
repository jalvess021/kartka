<template>
  <section class="bg-white dark:bg-gray-900">
    <div class="container px-6 py-8 mx-auto">
      <div class="lg:flex lg:-mx-2">
        <!-- Categorias -->
        <div class="space-y-3 lg:w-1/5 lg:px-2 lg:space-y-4">
          <a href="#" class="block font-medium text-gray-500 dark:text-gray-300 hover:underline">Novos</a>
          <a href="#" class="block font-medium text-gray-500 dark:text-gray-300 hover:underline">Usados</a>
          <a href="#" class="block font-medium text-blue-600 dark:text-blue-500 hover:underline">Em alta</a>
          <a href="#" class="block font-medium text-gray-500 dark:text-gray-300 hover:underline">Melhores avaliados</a>
        </div>

        <!-- Produtos -->
        <div class="mt-6 lg:mt-0 lg:px-2 lg:w-4/5">
          <div class="flex items-center justify-between text-sm tracking-widest uppercase">
            <p class="text-gray-500 dark:text-gray-300">{{products.length}} Items</p>
            <div class="flex items-center">
              <p class="text-gray-500 dark:text-gray-300">Classificar</p>
              <select class="font-medium text-gray-700 bg-transparent dark:text-gray-500 focus:outline-none">
                <option value="#">Recomendados</option>
                <option value="#">Tamanho</option>
                <option value="#">Preço</option>
              </select>
            </div>
          </div>

          <div class="grid grid-cols-1 gap-8 mt-8 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
            <div class="flex flex-col items-center justify-center w-full max-w-lg mx-auto" v-for="(product, index) in products" :key="index">
              <img :src="product.Image" alt="Product Image" class="object-cover w-full rounded-md h-72 xl:h-80" />
              <h4 class="mt-2 text-lg font-medium text-gray-700 dark:text-gray-200">{{ product.Name }}</h4>
              <p class="text-blue-500">R$ {{ product.Price }}</p>
              <button class="flex items-center justify-center w-full px-2 py-2 mt-4 font-medium tracking-wide text-white capitalize transition-colors duration-200 transform bg-gray-800 rounded-md hover:bg-gray-700 focus:outline-none focus:bg-gray-700">
                <svg xmlns="http://www.w3.org/2000/svg" class="w-5 h-5 mx-1" viewBox="0 0 20 20" fill="currentColor">
                  <path d="M3 1a1 1 0 000 2h1.22l.305 1.222a.997.997 0 00.01.042l1.358 5.43-.893.892C3.74 11.846 4.632 14 6.414 14H15a1 1 0 000-2H6.414l1-1H14a1 1 0 00.894-.553l3-6A1 1 0 0017 3H6.28l-.31-1.243A1 1 0 005 1H3zM16 16.5a1.5 1.5 0 11-3 0 1.5 1.5 0 013 0zM6.5 18a1.5 1.5 0 100-3 1.5 1.5 0 000 3z" />
                </svg>
                <span class="mx-1">Adicionar ao carrinho</span>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
  import { ref, onMounted } from 'vue';
  import axios from 'axios';

  interface Product {
    ID: string;
    Name: string;
    Price: number; 
    Image: string | null;
  }

  const products = ref<Product[]>([]);

  const loading = ref<boolean>(true);
  const error = ref<string | null>(null);

  const fetchProducts = async () => {
    try {
      const response = await axios.get('http://127.0.0.1:8282/api/products'); // URL da sua API
      products.value = response.data.map((product: Product) => ({
        ...product,
        Image: product.Image || "https://cdn.awsli.com.br/600x450/2530/2530043/produto/318397835/produto-teste_6355364765230299551-5dca5b60c67efea46e15917435354683-640-0-l0z22cg2ta.jpg",
      }));
      console.log(response)
    } catch (err) {
      error.value = 'Erro ao buscar produtos!';
      console.error(err);
    } finally {
      loading.value = false;
    }
  };
  
  onMounted(() => {
    fetchProducts();
  });
</script>
