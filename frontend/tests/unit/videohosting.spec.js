import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import videohosting from '@/components/videohosting.vue'

describe('videohosting.vue', () => {
  it('renders props.title when passed', () => {
    const title = 'Название'
    const wrapper = shallowMount(videohosting, {
      propsData: { title }
    })
    expect(wrapper.text()).to.include(title)
  })
  it('renders title', () => {
    const wrapper = shallowMount(videohosting, {})
    expect(wrapper.text()).to.include('Название')
    expect(wrapper.text()).to.include('Описание')
    expect(wrapper.text()).to.include('Ссылка на видео')
  })
})
