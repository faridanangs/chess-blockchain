package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faridanangs/checkers/x/checkers/types"
)

// SetTodo set a specific todo in the store from its index
func (k Keeper) SetTodo(ctx sdk.Context, todo types.Todo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TodoKeyPrefix))
	b := k.cdc.MustMarshal(&todo)
	store.Set(types.TodoKey(
		todo.Index,
	), b)
}

// GetTodo returns a todo from its index
func (k Keeper) GetTodo(
	ctx sdk.Context,
	index string,

) (val types.Todo, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TodoKeyPrefix))

	b := store.Get(types.TodoKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTodo removes a todo from the store
func (k Keeper) RemoveTodo(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TodoKeyPrefix))
	store.Delete(types.TodoKey(
		index,
	))
}

// GetAllTodo returns all todo
func (k Keeper) GetAllTodo(ctx sdk.Context) (list []types.Todo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TodoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Todo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
