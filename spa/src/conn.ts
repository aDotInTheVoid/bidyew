import { Store } from 'vuex';
import { State } from '@/store';

export function handleEvent(evt: any, store: Store<State>) {
    let message = JSON.parse(evt.data);
    let kind = message.Kind;
    let data = message.Data;

    store.commit(kind, data);
}