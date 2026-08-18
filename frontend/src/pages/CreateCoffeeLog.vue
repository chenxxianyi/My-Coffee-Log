<template>
  <div class="flex-1 w-full flex flex-col justify-between bg-coffee-warmWhite text-coffee-charcoal">
    
    <!-- Header -->
    <div class="px-5 py-3.5 border-b border-coffee-cream flex justify-between items-center bg-coffee-warmWhite/95 backdrop-blur-sm sticky top-0 z-10 select-none">
      <button
        @click="router.push('/home')"
        class="grid w-9 h-9 place-items-center -ml-2 rounded-full text-coffee-brown hover:bg-coffee-cream/60 hover:text-coffee-espresso transition-colors"
        aria-label="关闭并返回首页"
      >
        <X class="w-5 h-5" />
      </button>
      <div class="text-center leading-none">
        <h1 class="font-serif text-[17px] font-semibold tracking-wide text-coffee-espresso">
          {{ logMode === 'quick' ? '记录这一杯' : '精细记录' }}
        </h1>
        <p class="mt-1.5 text-[9px] tracking-[0.16em] text-coffee-softGray">
          {{ logMode === 'quick' ? '选好三项，即可保存' : `步骤 ${step} / 3` }}
        </p>
      </div>
      <button
        type="button"
        class="min-w-[64px] -mr-1 rounded-full border border-coffee-latte/35 bg-coffee-cream/35 px-2.5 py-1.5 text-[9px] font-semibold tracking-wider text-coffee-brown hover:border-coffee-brown/40 hover:bg-coffee-cream/65 transition-colors"
        @click="logMode = logMode === 'quick' ? 'detailed' : 'quick'"
      >
        {{ logMode === 'quick' ? '精细模式' : '快捷模式' }}
      </button>
    </div>

    <!-- Progress Indicator Bar (only in detailed mode) -->
    <div v-if="logMode === 'detailed'" class="h-1 bg-coffee-cream w-full flex select-none">
      <div class="h-full bg-coffee-brown transition-all duration-300" :style="{ width: (step * 33.3) + '%' }"></div>
    </div>

    <!-- Form Body -->
    <div class="flex-1 overflow-y-auto px-6 py-5">

      <!-- ==================== QUICK LOG MODE ==================== -->
      <div v-if="logMode === 'quick'" class="space-y-5">
        <div class="flex items-center justify-between rounded-sm border border-coffee-cream bg-coffee-cream/25 px-3.5 py-3 select-none">
          <div>
            <p class="font-serif text-sm font-semibold text-coffee-espresso">留下此刻，不必写得很满</p>
            <p class="mt-1 text-[9px] tracking-wider text-coffee-softGray">封面、类型和心情已预选，可直接保存</p>
          </div>
          <AppIcon name="coffee" :size="20" class="text-coffee-brown" />
        </div>

        <!-- Photo Selection -->
        <div class="space-y-2">
          <div class="quick-step-header">
            <span class="quick-step-number">01</span>
            <label class="quick-step-title">选择封面</label>
            <span class="quick-step-note">点选或拍照</span>
          </div>
          <div class="grid grid-cols-5 gap-2">
            <div 
              v-for="(imgUrl, idx) in store.DEFAULT_PHOTOS" 
              :key="idx"
              @click="quickForm.image_url = imgUrl"
              class="aspect-square relative cursor-pointer overflow-hidden rounded-sm border transition-all"
              :class="quickForm.image_url === imgUrl ? 'border-2 border-coffee-brown scale-[1.02]' : 'border-transparent opacity-80 hover:opacity-100'"
            >
              <img :src="imgUrl" class="w-full h-full object-cover">
              <div v-if="quickForm.image_url === imgUrl" class="absolute inset-0 bg-coffee-espresso/20 flex items-center justify-center text-white">
                <Check class="w-4 h-4" />
              </div>
            </div>

            <button
              type="button"
              class="aspect-square relative overflow-hidden rounded-sm border border-dashed border-coffee-latte/60 bg-coffee-cream/20 text-coffee-brown transition-colors hover:border-coffee-brown"
              :class="quickForm.image_url === uploadedImageUrl && uploadedImageUrl ? 'border-solid border-coffee-brown ring-1 ring-coffee-brown' : ''"
              aria-label="拍摄或上传封面"
              @click="triggerFileSelect"
            >
              <img v-if="uploadedImageUrl && !isUploading" :src="uploadedImageUrl" class="absolute inset-0 w-full h-full object-cover" alt="已上传封面">
              <span v-else class="absolute inset-0 flex flex-col items-center justify-center gap-1">
                <span v-if="isUploading" class="w-4 h-4 border-2 border-coffee-espresso border-t-transparent rounded-full animate-spin"></span>
                <Camera v-else class="w-4 h-4" />
                <span class="text-[8px] font-semibold">拍照</span>
              </span>
              <span v-if="quickForm.image_url === uploadedImageUrl && uploadedImageUrl && !isUploading" class="absolute inset-0 grid place-items-center bg-coffee-espresso/20 text-white">
                <Check class="w-4 h-4" />
              </span>
            </button>
          </div>

          <input 
            type="file" 
            ref="fileInput" 
            accept="image/*" 
            @change="handleFileChange" 
            class="hidden"
          >
        </div>

        <!-- Coffee Type Selection -->
        <div class="space-y-2">
          <div class="quick-step-header">
            <span class="quick-step-number">02</span>
            <label class="quick-step-title">咖啡类型</label>
            <span class="quick-step-note">单选</span>
          </div>
          <div class="grid grid-cols-3 gap-2">
            <button 
              v-for="t in typePresets" 
              :key="t.val"
              @click="quickForm.coffee_type = t.val"
              type="button"
              class="min-h-10 px-2 py-2 border rounded-sm text-center text-xs font-serif transition-all duration-200"
              :class="quickForm.coffee_type === t.val 
                ? 'bg-coffee-cream border-coffee-brown ring-1 ring-coffee-brown text-coffee-espresso' 
                : 'bg-coffee-cream/30 border-coffee-latte/50 hover:border-coffee-brown text-coffee-espresso'"
            >
              {{ t.shortLabel }}
            </button>
          </div>
        </div>

        <!-- Mood Selection -->
        <div class="space-y-2">
          <div class="quick-step-header">
            <span class="quick-step-number">03</span>
            <label class="quick-step-title">此时心情</label>
            <span class="quick-step-note">单选</span>
          </div>
          <div class="grid grid-cols-4 gap-2">
            <button
              v-for="m in moodPresets"
              :key="m.val"
              @click="quickForm.mood = m.val"
              type="button"
              class="min-h-[68px] p-2 border rounded-sm text-xs font-serif transition-all flex flex-col items-center justify-center gap-1"
              :class="quickForm.mood === m.val
                ? 'border-coffee-brown bg-coffee-cream text-coffee-espresso font-semibold'
                : 'border-coffee-latte/50 text-coffee-softGray hover:border-coffee-brown'"
            >
              <span class="icon-well" :class="quickForm.mood === m.val ? 'icon-well--active' : ''">
                <AppIcon :name="m.icon" :size="15" />
              </span>
              <span>{{ m.label }}</span>
            </button>
          </div>
        </div>

        <!-- Progressive disclosure for optional details -->
        <div class="pt-1 border-t border-coffee-cream/80">
          <button
            type="button"
            class="w-full flex items-center gap-3 rounded-sm px-1 py-3 text-left group"
            :aria-expanded="showQuickExtras"
            @click="showQuickExtras = !showQuickExtras"
          >
            <span class="grid w-8 h-8 flex-shrink-0 place-items-center rounded-full bg-coffee-cream/60 text-coffee-brown">
              <Plus v-if="!showQuickExtras" class="w-4 h-4" />
              <X v-else class="w-3.5 h-3.5" />
            </span>
            <span class="flex-1">
              <span class="block text-[11px] font-semibold tracking-wide text-coffee-espresso">添加更多细节</span>
              <span class="mt-1 block text-[9px] text-coffee-softGray">风味、生活标签、名称与 AI 文案</span>
            </span>
            <span v-if="quickExtrasCount" class="rounded-full bg-coffee-cream px-2 py-1 text-[8px] font-semibold text-coffee-brown">
              已添加 {{ quickExtrasCount }} 项
            </span>
            <ChevronDown class="w-4 h-4 text-coffee-softGray transition-transform duration-200" :class="showQuickExtras ? 'rotate-180' : ''" />
          </button>

          <Transition name="fade">
            <div v-if="showQuickExtras" class="mt-2 space-y-6 rounded-sm border border-coffee-cream/80 bg-white/20 p-4">
        <!-- Lifestyle Tags: Mood / Scene / Pairing -->
        <div class="space-y-3">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block select-none">生活标签 <span class="text-coffee-softGray font-normal normal-case">(选填)</span></label>
          <!-- Mood Tags -->
          <div class="space-y-1.5">
            <span class="text-[9px] uppercase tracking-wider text-coffee-softGray font-semibold select-none">心情 Mood</span>
            <div class="flex flex-wrap gap-1.5">
              <button
                v-for="t in LIFESTYLE_MOOD_TAGS"
                :key="t.val"
                @click="toggleLifestyleTag(quickForm.mood_tags, t.val)"
                type="button"
                class="lifestyle-chip"
                :class="quickForm.mood_tags.includes(t.val)
                  ? 'bg-amber-100 text-amber-800 border-amber-400'
                  : 'bg-coffee-cream/30 text-coffee-espresso border-coffee-latte/50 hover:border-coffee-brown'"
              >
                <AppIcon :name="t.icon" :size="12" />
                <span>{{ t.label }}</span>
              </button>
            </div>
          </div>
          <!-- Scene Tags -->
          <div class="space-y-1.5">
            <span class="text-[9px] uppercase tracking-wider text-coffee-softGray font-semibold select-none">场景 Scene</span>
            <div class="flex flex-wrap gap-1.5">
              <button
                v-for="t in LIFESTYLE_SCENE_TAGS"
                :key="t.val"
                @click="toggleLifestyleTag(quickForm.scene_tags, t.val)"
                type="button"
                class="lifestyle-chip"
                :class="quickForm.scene_tags.includes(t.val)
                  ? 'bg-sky-100 text-sky-800 border-sky-400'
                  : 'bg-coffee-cream/30 text-coffee-espresso border-coffee-latte/50 hover:border-coffee-brown'"
              >
                <AppIcon :name="t.icon" :size="12" />
                <span>{{ t.label }}</span>
              </button>
            </div>
          </div>
          <!-- Pairing Tags -->
          <div class="space-y-1.5">
            <span class="text-[9px] uppercase tracking-wider text-coffee-softGray font-semibold select-none">搭配 Pairing</span>
            <div class="flex flex-wrap gap-1.5">
              <button
                v-for="t in LIFESTYLE_PAIRING_TAGS"
                :key="t.val"
                @click="toggleLifestyleTag(quickForm.pairing_tags, t.val)"
                type="button"
                class="lifestyle-chip"
                :class="quickForm.pairing_tags.includes(t.val)
                  ? 'bg-rose-100 text-rose-800 border-rose-400'
                  : 'bg-coffee-cream/30 text-coffee-espresso border-coffee-latte/50 hover:border-coffee-brown'"
              >
                <AppIcon :name="t.icon" :size="12" />
                <span>{{ t.label }}</span>
              </button>
            </div>
          </div>
        </div>

        <!-- Flavor Impression Quick Selector -->
        <div class="space-y-3">
          <div class="flex justify-between items-center">
            <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso select-none">风味印象 <span class="text-coffee-softGray font-normal normal-case">(选填)</span></label>
            <button
              v-if="quickForm.flavor_preset"
              @click="quickForm.flavor_preset = ''"
              type="button"
              class="text-[9px] text-coffee-softGray hover:text-coffee-brown tracking-wider transition-colors"
            >清除选择</button>
          </div>

          <div class="grid grid-cols-2 gap-2">
            <button
              v-for="f in flavorPresets"
              :key="f.val"
              @click="quickForm.flavor_preset = quickForm.flavor_preset === f.val ? '' : f.val"
              type="button"
              class="p-3 border rounded-sm text-left transition-all duration-200 flex items-center gap-3"
              :class="quickForm.flavor_preset === f.val
                ? 'border-coffee-brown bg-coffee-cream text-coffee-espresso'
                : 'border-coffee-latte/50 text-coffee-softGray hover:border-coffee-brown'"
            >
              <span class="icon-well" :class="quickForm.flavor_preset === f.val ? 'icon-well--active' : ''">
                <AppIcon :name="f.icon" :size="17" />
              </span>
              <div>
                <div class="text-xs font-semibold font-serif leading-none">{{ f.label }}</div>
                <div class="text-[10px] mt-1 opacity-70">{{ f.desc }}</div>
              </div>
            </button>
          </div>

        </div>

        <!-- Optional: Coffee Name (quick fill) -->
        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block">咖啡名称 <span class="text-coffee-softGray font-normal normal-case">(选填)</span></label>
          <input 
            type="text" 
            v-model="quickForm.coffee_name" 
            placeholder="不填则自动生成" 
            class="w-full p-3 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm font-serif text-sm transition-colors"
          >
        </div>

        <label class="flex items-center gap-3 rounded-sm border border-coffee-cream bg-coffee-cream/25 p-3 cursor-pointer">
          <input
            v-model="generateAI"
            type="checkbox"
            class="h-4 w-4 flex-shrink-0 accent-coffee-espresso"
          >
          <span class="flex-1">
            <span class="block text-[10px] font-semibold tracking-wide text-coffee-espresso">生成 AI 感官文案</span>
            <span class="mt-1 block text-[9px] leading-relaxed text-coffee-softGray">关闭时仍会保存本地摘要</span>
          </span>
        </label>
            </div>
          </Transition>
        </div>
      </div>

      <!-- ==================== DETAILED LOG MODE ==================== -->

      <!-- STEP 1: Basic Information -->
      <div v-if="logMode === 'detailed' && step === 1" class="space-y-6">
        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block">1. 咖啡/豆子名称 *</label>
          <input 
            type="text" 
            v-model="form.coffee_name" 
            placeholder="例如: 埃塞俄比亚 耶加雪菲" 
            class="w-full p-3 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm font-serif text-sm transition-colors"
          >
        </div>

        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block select-none">2. 冲煮咖啡类型 *</label>
          <div class="grid grid-cols-2 gap-2">
            <button 
              v-for="t in typePresets" 
              :key="t.val"
              @click="form.coffee_type = t.val"
              type="button"
              class="p-3 border rounded-sm text-center text-xs font-serif font-light transition-all duration-200"
              :class="form.coffee_type === t.val 
                ? 'bg-coffee-cream border-coffee-brown ring-1 ring-coffee-brown text-coffee-espresso' 
                : 'bg-coffee-cream/30 border-coffee-latte/50 hover:border-coffee-brown text-coffee-espresso'"
            >
              {{ t.label }}
            </button>
          </div>
        </div>

        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block select-none">3. 手账封面图 (点击预设或上传本地照片)</label>
          <div class="grid grid-cols-4 gap-2">
            <div 
              v-for="(imgUrl, idx) in store.DEFAULT_PHOTOS" 
              :key="idx"
              @click="form.image_url = imgUrl"
              class="aspect-square relative cursor-pointer overflow-hidden rounded-sm border transition-all"
              :class="form.image_url === imgUrl ? 'border-2 border-coffee-brown scale-[1.02]' : 'border-transparent opacity-80 hover:opacity-100'"
            >
              <img :src="imgUrl" class="w-full h-full object-cover">
              <div v-if="form.image_url === imgUrl" class="absolute inset-0 bg-coffee-espresso/20 flex items-center justify-center text-white">
                <Check class="w-4 h-4" />
              </div>
            </div>
          </div>

          <!-- Upload Local Photo Component inside Create Journal -->
          <div 
            @click="triggerFileSelect"
            class="mt-3 p-3 border border-dashed border-coffee-latte/60 hover:border-coffee-brown bg-coffee-cream/15 rounded-sm cursor-pointer transition-colors flex items-center justify-center gap-2 select-none"
          >
            <template v-if="isUploading">
              <div class="w-4 h-4 border-2 border-coffee-espresso border-t-transparent rounded-full animate-spin"></div>
              <span class="text-xs text-coffee-espresso">正在上传相片...</span>
            </template>
            <template v-else-if="isLocalUploaded">
              <div class="w-6 h-6 rounded-full overflow-hidden border border-coffee-espresso">
                <img :src="form.image_url" class="w-full h-full object-cover">
              </div>
              <span class="text-xs text-green-700 font-medium">本地咖啡照片上传并设为封面！</span>
            </template>
            <template v-else>
              <Plus class="w-4 h-4 text-coffee-softGray" />
              <span class="text-xs text-coffee-espresso font-medium">使用手机拍摄/本地相片作为手账封面</span>
            </template>
          </div>
          <!-- Hidden Native File Input -->
          <input 
            type="file" 
            ref="fileInput" 
            accept="image/*" 
            @change="handleFileChange" 
            class="hidden"
          >
        </div>

        <!-- 4. Diary Writing Section -->
        <div class="space-y-2">
          <div class="flex items-center gap-2">
            <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso select-none">4. 手账日记 <span class="text-coffee-softGray font-normal normal-case">(选填)</span></label>
          </div>
          <div class="relative">
            <textarea 
              v-model="form.notes"
              rows="5"
              placeholder="今天在哪里？除了咖啡还有什么让你印象深刻？就用这段文字，把这一刻永久留下来……"
              class="w-full p-4 bg-coffee-cream/30 border border-coffee-latte/40 focus:border-coffee-brown focus:outline-none rounded-sm text-sm font-serif leading-relaxed resize-none transition-colors placeholder:text-coffee-softGray/60 placeholder:italic"
            ></textarea>
            <div class="absolute bottom-3 right-3 text-[9px] text-coffee-softGray/50 font-mono select-none">{{ form.notes.length }} 字</div>
          </div>
          <p class="text-[9px] text-coffee-softGray/70 italic select-none">AI 将在保存时参考这段文字，生成更有温度的感官评语。</p>
        </div>
      </div>

      <!-- STEP 2: Sensory Sliders with live SVG radar rendering -->
      <div v-if="logMode === 'detailed' && step === 2" class="space-y-6">
        <div class="flex justify-between items-center mb-2 select-none">
          <span class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso">感官风味指纹 / 6维风味参数</span>
          <span class="text-[10px] text-coffee-softGray italic">0 (无感) - 5 (浓郁)</span>
        </div>

        <!-- Horizontal layout: sliders on left, live responsive radar chart on right -->
        <div class="flex gap-4 items-center">
          <div class="flex-1 space-y-4">
            <!-- Render 6 dynamic sliders -->
            <div v-for="s in sliderSpecs" :key="s.key" class="space-y-1">
              <div class="flex justify-between text-xs text-coffee-espresso">
                <span>{{ s.label }}</span>
                <span class="font-semibold font-mono">{{ form[s.key] }}</span>
              </div>
              <input 
                type="range" 
                min="0" 
                max="5" 
                step="1" 
                v-model.number="form[s.key]" 
                class="w-full h-1 bg-coffee-cream rounded-lg appearance-none cursor-pointer accent-coffee-brown"
              >
            </div>
          </div>

          <!-- Dynamic SVG Radar Chart Component (Instant reactive updates!) -->
          <div class="w-[130px] h-[130px] flex-shrink-0 bg-coffee-cream/40 rounded-full flex items-center justify-center p-1 border border-coffee-latte/40 select-none">
            <FlavorRadarChart 
              :values="[form.acidity, form.bitterness, form.sweetness, form.body, form.aroma, form.aftertaste]"
              :size="120"
              :show-labels="false"
              :dot-radius="2.0"
            />
          </div>
        </div>

        <!-- Brew Parameters (optional, for advanced users) -->
        <div class="space-y-3 mt-4 pt-4 border-t border-coffee-cream">
          <div class="flex items-center gap-2 select-none">
            <span class="text-[9px] uppercase tracking-[0.2em] font-semibold text-coffee-softGray">冲煮参数 / Brew Params</span>
            <span class="text-[9px] text-coffee-softGray/60 normal-case">(选填)</span>
          </div>
          <div class="grid grid-cols-3 gap-3">
            <div class="space-y-1">
              <label class="text-[9px] uppercase tracking-wider text-coffee-softGray select-none">粉水比</label>
              <input type="text" v-model="form.brew_ratio" placeholder="1:15" class="w-full p-2 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-xs">
            </div>
            <div class="space-y-1">
              <label class="text-[9px] uppercase tracking-wider text-coffee-softGray select-none">水温</label>
              <input type="text" v-model="form.water_temp" placeholder="92°C" class="w-full p-2 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-xs">
            </div>
            <div class="space-y-1">
              <label class="text-[9px] uppercase tracking-wider text-coffee-softGray select-none">研磨度</label>
              <input type="text" v-model="form.grind_size" placeholder="中细" class="w-full p-2 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-xs">
            </div>
          </div>
        </div>
      </div>

      <!-- STEP 3: Flavor tags, Mood, Spot & Notes -->
      <div v-if="logMode === 'detailed' && step === 3" class="space-y-5">
        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block font-medium select-none">1. 风味特征标签 (点击多选)</label>
          <div class="flex flex-wrap gap-1.5">
            <button
              v-for="tag in tagPresets"
              :key="tag.name"
              @click="toggleTag(tag.name)"
              type="button"
              class="px-3 py-1 text-[11px] border rounded-full transition-all duration-150"
              :class="form.flavor_tags.includes(tag.name)
                ? 'bg-coffee-espresso text-coffee-warmWhite border-coffee-espresso'
                : 'bg-coffee-cream/40 text-coffee-espresso border-coffee-latte/50 hover:border-coffee-brown'"
            >
              {{ tag.label }}
            </button>
          </div>
        </div>

        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block font-medium select-none">2. 此时此地心情 / Mood</label>
          <div class="grid grid-cols-4 gap-2">
            <button
              v-for="m in moodPresets"
              :key="m.val"
              @click="form.mood = m.val"
              type="button"
              class="p-2 border rounded-sm text-xs font-serif transition-all flex flex-col items-center gap-1"
              :class="form.mood === m.val
                ? 'border-coffee-brown bg-coffee-cream text-coffee-espresso font-semibold'
                : 'border-coffee-latte/50 text-coffee-softGray hover:border-coffee-brown'"
            >
              <span class="icon-well" :class="form.mood === m.val ? 'icon-well--active' : ''">
                <AppIcon :name="m.icon" :size="15" />
              </span>
              <span>{{ m.label }}</span>
            </button>
          </div>
        </div>

        <div class="space-y-2 relative">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block">3. 咖啡出品馆 / Shop & Spot</label>
          <input 
            type="text" 
            v-model="form.shop_name" 
            @focus="loadShopNames"
            @input="filterShopNames"
            placeholder="例如: Blue Bottle, 上海" 
            class="w-full p-2.5 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-xs"
          >
          <!-- Autocomplete Dropdown -->
          <div v-if="showShopDropdown && filteredShopNames.length > 0" class="absolute left-0 right-0 top-full mt-1 bg-coffee-warmWhite border border-coffee-latte/50 rounded-sm shadow-lg z-20 max-h-32 overflow-y-auto">
            <button 
              v-for="name in filteredShopNames" 
              :key="name"
              @mousedown.prevent="selectShopName(name)"
              class="w-full text-left px-3 py-2 text-xs text-coffee-espresso hover:bg-coffee-cream/60 transition-colors flex items-center gap-2"
            >
              <MapPin class="w-3 h-3 text-coffee-softGray flex-shrink-0" />
              <span>{{ name }}</span>
            </button>
          </div>
        </div>

        <!-- Coffee Bean Selection (optional, for advanced users) -->
        <div class="space-y-2">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block select-none">4. 咖啡豆档案 <span class="text-coffee-softGray font-normal normal-case">(选填)</span></label>
          <!-- Select existing bean -->
          <div class="relative">
            <button 
              @click="loadBeanList" 
              type="button"
              class="w-full p-2.5 bg-coffee-cream/40 border border-coffee-latte/60 rounded-sm text-xs text-left flex items-center justify-between hover:border-coffee-brown transition-colors"
            >
              <span :class="form.bean_id ? 'text-coffee-espresso font-medium' : 'text-coffee-softGray'">
                {{ selectedBeanLabel || '选择已保存的咖啡豆...' }}
              </span>
              <ChevronDown class="w-3.5 h-3.5 text-coffee-softGray" />
            </button>
            <!-- Bean dropdown -->
            <div v-if="showBeanDropdown && beanList.length > 0" class="absolute left-0 right-0 top-full mt-1 bg-coffee-warmWhite border border-coffee-latte/50 rounded-sm shadow-lg z-20 max-h-36 overflow-y-auto">
              <button 
                v-for="bean in beanList" 
                :key="bean.id"
                @mousedown.prevent="selectBean(bean)"
                class="w-full text-left px-3 py-2 text-xs text-coffee-espresso hover:bg-coffee-cream/60 transition-colors flex items-center justify-between"
              >
                <div>
                  <span class="font-medium">{{ bean.name }}</span>
                  <span v-if="bean.origin" class="text-coffee-softGray ml-1.5">{{ bean.origin }}</span>
                </div>
                <span v-if="bean.roast_level" class="text-[9px] text-coffee-brown">{{ bean.roast_level }}</span>
              </button>
            </div>
          </div>
          <!-- Or fill inline -->
          <div class="space-y-2 pt-1">
            <div class="text-[9px] text-coffee-softGray select-none">或手动填写豆子信息</div>
            <div class="grid grid-cols-2 gap-2">
              <input type="text" v-model="form.bean_name" placeholder="豆子名称" class="p-2 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-xs">
              <input type="text" v-model="beanOrigin" placeholder="产地 (如: 埃塞俄比亚)" class="p-2 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-xs">
            </div>
            <div class="grid grid-cols-3 gap-2">
              <input type="text" v-model="beanProcess" placeholder="处理法" class="p-2 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-xs">
              <input type="text" v-model="beanRoast" placeholder="烘焙度" class="p-2 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-xs">
              <input type="text" v-model="beanRoaster" placeholder="烘焙商" class="p-2 bg-coffee-cream/40 border border-coffee-latte/60 focus:border-coffee-brown focus:outline-none rounded-sm text-xs">
            </div>
          </div>
        </div>

        <!-- Lifestyle Tags: Mood / Scene / Pairing -->
        <div class="space-y-3">
          <label class="text-[10px] uppercase tracking-[0.2em] font-semibold text-coffee-espresso block font-medium select-none">5. 生活标签 <span class="text-coffee-softGray font-normal normal-case">(选填)</span></label>
          <!-- Mood Tags -->
          <div class="space-y-1.5">
            <span class="text-[9px] uppercase tracking-wider text-coffee-softGray font-semibold select-none">心情 Mood</span>
            <div class="flex flex-wrap gap-1.5">
              <button
                v-for="t in LIFESTYLE_MOOD_TAGS"
                :key="t.val"
                @click="toggleLifestyleTag(form.mood_tags!, t.val)"
                type="button"
                class="lifestyle-chip"
                :class="form.mood_tags?.includes(t.val)
                  ? 'bg-amber-100 text-amber-800 border-amber-400'
                  : 'bg-coffee-cream/30 text-coffee-espresso border-coffee-latte/50 hover:border-coffee-brown'"
              >
                <AppIcon :name="t.icon" :size="12" />
                <span>{{ t.label }}</span>
              </button>
            </div>
          </div>
          <!-- Scene Tags -->
          <div class="space-y-1.5">
            <span class="text-[9px] uppercase tracking-wider text-coffee-softGray font-semibold select-none">场景 Scene</span>
            <div class="flex flex-wrap gap-1.5">
              <button
                v-for="t in LIFESTYLE_SCENE_TAGS"
                :key="t.val"
                @click="toggleLifestyleTag(form.scene_tags!, t.val)"
                type="button"
                class="lifestyle-chip"
                :class="form.scene_tags?.includes(t.val)
                  ? 'bg-sky-100 text-sky-800 border-sky-400'
                  : 'bg-coffee-cream/30 text-coffee-espresso border-coffee-latte/50 hover:border-coffee-brown'"
              >
                <AppIcon :name="t.icon" :size="12" />
                <span>{{ t.label }}</span>
              </button>
            </div>
          </div>
          <!-- Pairing Tags -->
          <div class="space-y-1.5">
            <span class="text-[9px] uppercase tracking-wider text-coffee-softGray font-semibold select-none">搭配 Pairing</span>
            <div class="flex flex-wrap gap-1.5">
              <button
                v-for="t in LIFESTYLE_PAIRING_TAGS"
                :key="t.val"
                @click="toggleLifestyleTag(form.pairing_tags!, t.val)"
                type="button"
                class="lifestyle-chip"
                :class="form.pairing_tags?.includes(t.val)
                  ? 'bg-rose-100 text-rose-800 border-rose-400'
                  : 'bg-coffee-cream/30 text-coffee-espresso border-coffee-latte/50 hover:border-coffee-brown'"
              >
                <AppIcon :name="t.icon" :size="12" />
                <span>{{ t.label }}</span>
              </button>
            </div>
          </div>
        </div>

      </div>

    </div>

    <!-- Bottom Controls -->
    <div
      class="border-t border-coffee-cream bg-coffee-warmWhite/95 backdrop-blur-sm sticky bottom-0 z-10 select-none space-y-3"
      :class="logMode === 'quick' ? 'px-5 py-4' : 'p-6'"
    >
      <label v-if="logMode === 'detailed'" class="flex items-start gap-3 p-3 rounded-sm border border-coffee-cream bg-coffee-cream/25 cursor-pointer">
        <input
          v-model="generateAI"
          type="checkbox"
          class="mt-0.5 h-4 w-4 accent-coffee-espresso"
        >
        <span class="space-y-0.5">
          <span class="block text-[10px] uppercase tracking-[0.22em] font-semibold text-coffee-espresso">云端 AI 文案</span>
          <span class="block text-[10px] leading-relaxed text-coffee-softGray">开启后会将咖啡名称、风味分数与笔记发送至已配置的 AI 服务；关闭时仅使用本地摘要。</span>
        </span>
      </label>
      <div v-if="logMode === 'quick'" class="flex items-center justify-between text-[9px] tracking-wide">
        <span class="inline-flex items-center gap-1.5 font-semibold text-coffee-brown">
          <Check class="w-3 h-3" />
          {{ typePresets.find(t => t.val === quickForm.coffee_type)?.shortLabel }} · {{ moodPresets.find(m => m.val === quickForm.mood)?.label }}
        </span>
        <span class="text-coffee-softGray">信息已齐，可以直接保存</span>
      </div>
      <div class="flex gap-3">
        <!-- Detailed mode: show back button -->
        <button 
          v-if="logMode === 'detailed' && step > 1" 
          @click="step--" 
          class="flex-1 py-3 text-xs uppercase tracking-wider font-semibold bg-coffee-cream text-coffee-espresso hover:bg-coffee-latte transition-all rounded-sm"
        >
          上一步
        </button>
        <button 
          @click="handleNext" 
          class="flex-1 py-3 text-xs uppercase tracking-wider font-semibold bg-coffee-espresso text-coffee-warmWhite hover:bg-coffee-brown transition-all rounded-sm flex items-center justify-center gap-1.5"
          :disabled="isSubmitting"
        >
          <template v-if="isSubmitting">
            <div class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
            <span>{{ generateAI ? '正在保存并准备云端文案...' : '正在保存本地摘要...' }}</span>
          </template>
          <template v-else>
            <span v-if="logMode === 'quick'">保存这杯咖啡</span>
            <span v-else>{{ step === 3 ? (generateAI ? '保存并启用 AI 文案' : '保存手账') : '下一步' }}</span>
          </template>
        </button>
      </div>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useCoffeeLogStore, NewCoffeeLog } from '@/stores/coffeeLog'
import FlavorRadarChart from '@/components/charts/FlavorRadarChart.vue'
import AppIcon from '@/components/AppIcon.vue'
import request from '@/api/request'
import { X, Check, Plus, Camera, MapPin, ChevronDown } from 'lucide-vue-next'
import { LIFESTYLE_MOOD_TAGS, LIFESTYLE_SCENE_TAGS, LIFESTYLE_PAIRING_TAGS } from '@/constants/coffee'
import * as shopApi from '@/api/coffeeShop'
import * as beanApi from '@/api/coffeeBean'

const router = useRouter()
const route = useRoute()
const store = useCoffeeLogStore()

// Mode: 'quick' = single page, 'detailed' = 3-step wizard
const logMode = ref<'quick' | 'detailed'>('quick')
const step = ref(1)
const isSubmitting = ref(false)
const generateAI = ref(false)
const showQuickExtras = ref(false)

// File upload states
const fileInput = ref<HTMLInputElement | null>(null)
const isUploading = ref(false)
const isLocalUploaded = ref(false)
const uploadedImageUrl = ref('')

// Shop autocomplete
const shopNames = ref<string[]>([])
const showShopDropdown = ref(false)
const filteredShopNames = computed(() => {
  const q = form.shop_name.trim().toLowerCase()
  if (!q) return shopNames.value.slice(0, 8)
  return shopNames.value.filter(n => n.toLowerCase().includes(q)).slice(0, 8)
})

async function loadShopNames() {
  if (shopNames.value.length === 0) {
    try {
      shopNames.value = await shopApi.getShopNames()
    } catch { /* ignore */ }
  }
  showShopDropdown.value = true
}

function filterShopNames() {
  showShopDropdown.value = true
}

function selectShopName(name: string) {
  form.shop_name = name
  showShopDropdown.value = false
}

// Bean autocomplete
const beanList = ref<beanApi.CoffeeBean[]>([])
const showBeanDropdown = ref(false)
const beanOrigin = ref('')
const beanProcess = ref('')
const beanRoast = ref('')
const beanRoaster = ref('')
const selectedBeanLabel = computed(() => {
  if (!form.bean_id) return ''
  const found = beanList.value.find(b => b.id === form.bean_id)
  return found ? `${found.name}${found.origin ? ' - ' + found.origin : ''}` : ''
})

async function loadBeanList() {
  if (beanList.value.length === 0) {
    try {
      beanList.value = await beanApi.getBeanList()
    } catch { /* ignore */ }
  }
  showBeanDropdown.value = true
}

function selectBean(bean: beanApi.CoffeeBean) {
  form.bean_id = bean.id
  form.bean_name = ''
  beanOrigin.value = ''
  beanProcess.value = ''
  beanRoast.value = ''
  beanRoaster.value = ''
  showBeanDropdown.value = false
}

// Quick Log Form — minimal fields, defaults for the rest
const quickForm = reactive({
  image_url: store.DEFAULT_PHOTOS[1],
  coffee_type: 'Pour Over',
  mood: 'Calm',
  coffee_name: '',
  flavor_preset: '',
  mood_tags: [] as string[],
  scene_tags: [] as string[],
  pairing_tags: [] as string[]
})

const quickExtrasCount = computed(() => {
  return quickForm.mood_tags.length +
    quickForm.scene_tags.length +
    quickForm.pairing_tags.length +
    (quickForm.flavor_preset ? 1 : 0) +
    (quickForm.coffee_name.trim() ? 1 : 0) +
    (generateAI.value ? 1 : 0)
})

// Prefill from existing log (Brew Again)
onMounted(async () => {
  const fromLogId = route.query.from_log_id
  if (fromLogId) {
    const id = Number(fromLogId)
    if (!Number.isNaN(id)) {
      let sourceLog = store.getLogById(id)
      if (!sourceLog) {
        try { sourceLog = await store.fetchLogById(id) } catch { /* ignore */ }
      }
      if (sourceLog) {
        // Prefill quick form
        quickForm.coffee_type = sourceLog.coffee_type
        quickForm.mood = sourceLog.mood || 'Calm'
        quickForm.image_url = sourceLog.image_url || store.DEFAULT_PHOTOS[1]
        quickForm.mood_tags = [...(sourceLog.mood_tags || [])]
        quickForm.scene_tags = [...(sourceLog.scene_tags || [])]
        quickForm.pairing_tags = [...(sourceLog.pairing_tags || [])]
        showQuickExtras.value = quickForm.mood_tags.length > 0 || quickForm.scene_tags.length > 0 || quickForm.pairing_tags.length > 0
        if (sourceLog.image_url && !store.DEFAULT_PHOTOS.includes(sourceLog.image_url)) {
          uploadedImageUrl.value = sourceLog.image_url
        }
        // Prefill detailed form
        form.coffee_type = sourceLog.coffee_type
        form.image_url = sourceLog.image_url || store.DEFAULT_PHOTOS[1]
        form.acidity = sourceLog.acidity
        form.bitterness = sourceLog.bitterness
        form.sweetness = sourceLog.sweetness
        form.body = sourceLog.body
        form.aroma = sourceLog.aroma
        form.aftertaste = sourceLog.aftertaste
        form.flavor_tags = [...(sourceLog.flavor_tags || [])]
        form.mood = sourceLog.mood || 'Calm'
        form.shop_name = sourceLog.shop_name || ''
        form.notes = sourceLog.notes || ''
        form.mood_tags = [...(sourceLog.mood_tags || [])]
        form.scene_tags = [...(sourceLog.scene_tags || [])]
        form.pairing_tags = [...(sourceLog.pairing_tags || [])]
      }
    }
  }
})

// Detailed Log Form — full 3-step wizard
const form = reactive<NewCoffeeLog>({
  coffee_name: '',
  coffee_type: 'Pour Over',
  image_url: store.DEFAULT_PHOTOS[1],
  acidity: 4,
  bitterness: 1,
  sweetness: 3,
  body: 2,
  aroma: 5,
  aftertaste: 4,
  flavor_tags: ['citrus', 'floral'],
  mood: 'Calm',
  shop_name: '',
  notes: '',
  mood_tags: [],
  scene_tags: [],
  pairing_tags: [],
  bean_id: null as number | null,
  bean_name: '',
  brew_ratio: '',
  water_temp: '',
  grind_size: ''
})

// Auto-generate coffee name for quick log
const generateQuickName = (type: string) => {
  const typeMap: Record<string, string> = {
    'Pour Over': '手冲咖啡',
    'Latte': '拿铁',
    'Americano': '美式咖啡',
    'Cold Brew': '冷萃咖啡',
    'Espresso': '浓缩咖啡',
    'Dirty': 'Dirty 咖啡'
  }
  return typeMap[type] || '咖啡'
}

// Presets Specs
const typePresets = [
  { val: 'Pour Over', label: 'Pour Over / 手冲', shortLabel: '手冲' },
  { val: 'Latte', label: 'Latte / 拿铁', shortLabel: '拿铁' },
  { val: 'Americano', label: 'Americano / 美式', shortLabel: '美式' },
  { val: 'Cold Brew', label: 'Cold Brew / 冷萃', shortLabel: '冷萃' },
  { val: 'Espresso', label: 'Espresso / 浓缩', shortLabel: '浓缩' },
  { val: 'Dirty', label: 'Dirty / 脏咖啡', shortLabel: '脏咖啡' }
]

const sliderSpecs = [
  { key: 'acidity', label: '酸度 Acidity' },
  { key: 'bitterness', label: '苦感 Bitterness' },
  { key: 'sweetness', label: '甜感 Sweetness' },
  { key: 'body', label: '醇厚度 Body' },
  { key: 'aroma', label: '香气 Aroma' },
  { key: 'aftertaste', label: '余韵 Aftertaste' }
] as const

const tagPresets = [
  { name: 'floral', label: '花香' },
  { name: 'citrus', label: '柑橘' },
  { name: 'berry', label: '莓果' },
  { name: 'nutty', label: '坚果' },
  { name: 'chocolate', label: '巧克力' },
  { name: 'caramel', label: '焦糖' },
  { name: 'creamy', label: '奶油' },
  { name: 'winey', label: '酒香' }
]
const flavorPresets = [
  {
    val: 'bright',
    label: '清新明亮',
    desc: '高酸 · 轻盈',
    icon: 'bright',
    values: { acidity: 5, bitterness: 1, sweetness: 3, body: 2, aroma: 4, aftertaste: 3 }
  },
  {
    val: 'floral',
    label: '花果芬芳',
    desc: '果香 · 花香',
    icon: 'floral',
    values: { acidity: 4, bitterness: 1, sweetness: 4, body: 2, aroma: 5, aftertaste: 3 }
  },
  {
    val: 'smooth',
    label: '甜美柔滑',
    desc: '低酸 · 甜润',
    icon: 'smooth',
    values: { acidity: 2, bitterness: 1, sweetness: 5, body: 3, aroma: 3, aftertaste: 4 }
  },
  {
    val: 'bold',
    label: '浓郁醇厚',
    desc: '厚重 · 回甘',
    icon: 'bold',
    values: { acidity: 1, bitterness: 4, sweetness: 2, body: 5, aroma: 4, aftertaste: 5 }
  }
]

const moodPresets = [
  { val: 'Calm', label: '平静', icon: 'calm' },
  { val: 'Energetic', label: '愉悦', icon: 'energetic' },
  { val: 'Reflective', label: '沉浸', icon: 'reflective' },
  { val: 'Tired', label: '疲惫', icon: 'tired' }
]

const toggleTag = (tag: string) => {
  if (form.flavor_tags.includes(tag)) {
    form.flavor_tags = form.flavor_tags.filter((t: string) => t !== tag)
  } else {
    form.flavor_tags.push(tag)
  }
}

const toggleLifestyleTag = (arr: string[], tag: string) => {
  const idx = arr.indexOf(tag)
  if (idx >= 0) {
    arr.splice(idx, 1)
  } else {
    arr.push(tag)
  }
}

const handleNext = async () => {
  // Quick Log: save immediately
  if (logMode.value === 'quick') {
    isSubmitting.value = true
    const quickLog: NewCoffeeLog = {
      coffee_name: quickForm.coffee_name.trim() || generateQuickName(quickForm.coffee_type),
      coffee_type: quickForm.coffee_type,
      image_url: quickForm.image_url,
      mood: quickForm.mood,
      shop_name: 'Local Coffee Spot',
      notes: '一杯温润安静的手账记录。',
      generate_ai: generateAI.value,
      ...(flavorPresets.find(f => f.val === quickForm.flavor_preset)?.values ?? { acidity: 3, bitterness: 2, sweetness: 3, body: 3, aroma: 3, aftertaste: 3 }),
      flavor_tags: [],
      mood_tags: quickForm.mood_tags,
      scene_tags: quickForm.scene_tags,
      pairing_tags: quickForm.pairing_tags
    }
    try {
      const created = await store.addLog(quickLog)
      isSubmitting.value = false
      store.fetchStats()
      router.push(`/coffee/${created.id}?just_created=true`)
    } catch (e: any) {
      isSubmitting.value = false
      alert(e.message || '保存失败，请稍后重试')
    }
    return
  }

  // Detailed Log: step-by-step wizard
  if (step.value === 1) {
    if (!form.coffee_name.trim()) {
      alert('请输入咖啡名称哦')
      return
    }
    step.value = 2
  } 
  else if (step.value === 2) {
    step.value = 3
  } 
  else if (step.value === 3) {
    isSubmitting.value = true

    // If inline bean info provided and no bean_id selected, create bean first
    if (!form.bean_id && form.bean_name?.trim()) {
      try {
        const newBean = await beanApi.createCoffeeBean({
          name: form.bean_name!.trim(),
          origin: beanOrigin.value.trim() || undefined,
          processing_method: beanProcess.value.trim() || undefined,
          roast_level: beanRoast.value.trim() || undefined,
          roaster: beanRoaster.value.trim() || undefined
        })
        form.bean_id = newBean.id
      } catch (e: any) {
        // Bean creation failed (e.g. duplicate name) — try to find existing
        try {
          const beans = await beanApi.getBeanList()
          const existing = beans.find(b => b.name === form.bean_name!.trim())
          if (existing) form.bean_id = existing.id
        } catch { /* proceed without bean_id */ }
      }
    }

    const savedLog = {
      ...form,
      shop_name: form.shop_name.trim() || 'Local Coffee Spot',
      notes: form.notes.trim() || '一杯温润安静的手账记录。',
      generate_ai: generateAI.value
    }
    try {
      const created = await store.addLog(savedLog)
      isSubmitting.value = false
      store.fetchStats()
      router.push(`/coffee/${created.id}?just_created=true`)
    } catch (e: any) {
      isSubmitting.value = false
      alert(e.message || '保存失败，请稍后重试')
    }
  }
}

// Local File Upload Handlers
const triggerFileSelect = () => {
  if (fileInput.value) {
    fileInput.value.click()
  }
}

const handleFileChange = async (event: Event) => {
  const target = event.target as HTMLInputElement
  const files = target.files
  if (!files || files.length === 0) return

  const file = files[0]
  if (file.size > 5 * 1024 * 1024) {
    alert('照片大小不能超过 5MB 喔！')
    return
  }

  isUploading.value = true
  const formData = new FormData()
  formData.append('file', file)

  try {
    const res = await request.post('/upload', formData, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    }) as any

    if (res && res.url) {
      uploadedImageUrl.value = res.url
      // Set image on the active form based on current mode
      if (logMode.value === 'quick') {
        quickForm.image_url = res.url
      } else {
        form.image_url = res.url
      }
      isLocalUploaded.value = true
    } else {
      throw new Error('未获取到返回的图片地址')
    }
  } catch (e: any) {
    alert(e.message || '相片上传失败，请重试')
  } finally {
    isUploading.value = false
    target.value = '' // Clear input
  }
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active { transition: opacity 0.2s ease, transform 0.2s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; transform: translateY(-4px); }

.quick-step-header {
  display: flex;
  align-items: center;
  gap: 0.55rem;
  min-height: 1.5rem;
  user-select: none;
}

.quick-step-number {
  display: grid;
  width: 1.4rem;
  height: 1.4rem;
  place-items: center;
  border: 1px solid rgba(192, 160, 124, 0.45);
  border-radius: 9999px;
  color: #9c7b59;
  font-family: "Cormorant Garamond", serif;
  font-size: 10px;
  font-weight: 600;
}

.quick-step-title {
  color: #5c3d2e;
  font-size: 10px;
  font-weight: 700;
  letter-spacing: 0.14em;
}

.quick-step-note {
  margin-left: auto;
  color: #c0a07c;
  font-size: 9px;
  letter-spacing: 0.06em;
}

/* Range slider styling */
input[type="range"]::-webkit-slider-thumb {
  border-radius: 50%;
  width: 12px;
  height: 12px;
  background: #7A5638;
}

.lifestyle-chip {
  display: inline-flex;
  align-items: center;
  gap: 0.35rem;
  min-height: 1.75rem;
  padding: 0.25rem 0.65rem;
  border-width: 1px;
  border-radius: 9999px;
  font-size: 10px;
  line-height: 1;
  transition: color 150ms ease, border-color 150ms ease, background-color 150ms ease, transform 150ms ease;
}

.lifestyle-chip:active {
  transform: scale(0.97);
}

.icon-well {
  display: inline-grid;
  width: 1.75rem;
  height: 1.75rem;
  flex-shrink: 0;
  place-items: center;
  border: 1px solid rgba(192, 160, 124, 0.35);
  border-radius: 9999px;
  background: rgba(253, 232, 194, 0.35);
  color: #9c7b59;
  transition: color 180ms ease, border-color 180ms ease, background-color 180ms ease;
}

.icon-well--active {
  border-color: rgba(92, 61, 46, 0.32);
  background: rgba(92, 61, 46, 0.09);
  color: #5c3d2e;
}
</style>
