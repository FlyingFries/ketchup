//Package permissions provides simple yet powerful interface to managing permissions
// author: Karol Piłat <karol@vitresoft.com>
package permissions

import "strings"
import "fmt"

//Permissible is the interfece that wraps basic HasPermission method
type Permissible interface {
	HasPermission(perm string) bool
}

//Set is basic set of permissions
type Set map[string]bool

//HasPermission implements checking if permission is in Set
func (ps Set) HasPermission(perm string) bool {
	for {
		if v, ok := ps[perm]; ok {
			return v
		}
		li := strings.LastIndex(perm, ".")
		if li < 0 {
			return false
		}
		perm = perm[:li]
	}
}

//SetPermission sets permission
func (ps Set) SetPermission(perm string, value bool) {
	ps[perm] = value
}

//Apply can rewrite whole set of permissions
func (ps Set) Apply(perms Set, root string, override bool) error {
	for k, v := range perms {
		if !strings.HasPrefix(k, root+".") && k != root && root != "" {
			return fmt.Errorf("capability %s outside of root %s", k, root)
		}
		if old, has := ps[k]; !override && has && old != v {
			return fmt.Errorf("conflict on %s capability! %t != %t", k, old, v)
		}
		ps[k] = v
	}
	return nil
}
