import { FormGroupInput, Card, DropDown, Button } from '../components/index';
import { DatePicker } from 'element-ui';

/**
 * You can register global components here and use them as a plugin in your main Vue instance
 */

const GlobalComponents = {
  install(Vue) {
    Vue.component('fg-input', FormGroupInput);
    Vue.component('drop-down', DropDown);
    Vue.component('card', Card);
    Vue.component('p-button', Button);
    Vue.component(DatePicker.name, DatePicker);
  }
};

export default GlobalComponents;
