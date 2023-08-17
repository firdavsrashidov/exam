package controller

import (
	"app/models"
	"sort"
)

// 1-task
func (c *Controller) Sort(req *models.OrderGetListRequest) (*models.OrderGetList, error) {
	var resp = &models.OrderGetList{}
	var orderDateFilter []*models.Order
	getorder, err := c.OrderGetList(req)
	if err != nil {
		return nil, err
	}
	for _, ord := range getorder.Orders {
		orderDateFilter = append(orderDateFilter, ord)

	}
	sort.Slice(orderDateFilter, func(i, j int) bool {
		return orderDateFilter[i].DateTime > orderDateFilter[j].DateTime
	})
	resp.Count = len(orderDateFilter)
	resp.Orders = orderDateFilter
	return resp, nil
}

// 2-task
func (c *Controller) Filter(req *models.OrderGetListRequest) ([]*models.Order, error) {
	var orderDateFilter []*models.Order
	getorder, err := c.OrderGetList(req)
	if err != nil {
		return nil, err
	}
	for _, ord := range getorder.Orders {
		if ord.DateTime >= req.FromTime && ord.DateTime < req.ToTime {
			orderDateFilter = append(orderDateFilter, ord)
		}
	}

	return orderDateFilter, nil
}

// 4-task
func (c *Controller) UserCash(req *models.UserPrimaryKey) (map[string]int, error) {
	user := make(map[string]int)

	getorder, err := c.OrderGetList(&models.OrderGetListRequest{})
	if err != nil {
		return nil, err
	}

	getuser, err := c.UserGetById(req)

	for _, value := range getorder.Orders {
		if value.UserId == req.Id {
			if value.Status == true {
				getproduct, err := c.GetByIdPoduct(&models.ProductPrimaryKey{Id: value.Id})
				if err != nil {
					return nil, err
				}
				user[getuser.FirstName] += value.SumCount * getproduct.Price
			}
		}
	}
	return user, nil
}

// 5-task
func (c *Controller) ProductCountSold() (map[string]int, error) {
	product := make(map[string]int)

	getorder, err := c.OrderGetList(&models.OrderGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getorder.Orders {
		getproduct, err := c.GetByIdPoduct(&models.ProductPrimaryKey{Id: value.Id})
		if err != nil {
			return nil, err
		}
		if value.Status == true {
			product[getproduct.Name] += value.SumCount
		}

	}
	return product, nil
}

// 6-task
func (c *Controller) TopProducts() ([]*models.ProductsHistory, error) {
	var (
		prodctsMap = make(map[string]int)
		products   []*models.ProductsHistory
	)

	getOrder, err := c.OrderGetList(&models.OrderGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getOrder.Orders {
		getProduct, err := c.GetByIdPoduct(&models.ProductPrimaryKey{Id: value.Id})
		if err != nil {
			return nil, err
		}
		if value.Status == true {
			prodctsMap[getProduct.Name] += value.SumCount
		}
	}
	for k, v := range prodctsMap {
		products = append(products, &models.ProductsHistory{
			Name:  k,
			Count: v,
		})
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Count > products[j].Count
	})

	return products, nil
}

// 7-task
func (c *Controller) FailureProducts() ([]*models.ProductsHistory, error) {
	var (
		prodctsMap = make(map[string]int)
		products   []*models.ProductsHistory
	)

	getOrder, err := c.OrderGetList(&models.OrderGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getOrder.Orders {
		getProduct, err := c.GetByIdPoduct(&models.ProductPrimaryKey{Id: value.Id})
		if err != nil {
			return nil, err
		}
		if value.Status == true {
			prodctsMap[getProduct.Name] += value.SumCount
		}
	}
	for k, v := range prodctsMap {
		products = append(products, &models.ProductsHistory{
			Name:  k,
			Count: v,
		})
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Count < products[j].Count
	})

	return products, nil
}

// 8-task
func (c *Controller) TopTime() ([]*models.DateHistory, error) {
	var (
		toptimes = make(map[string]int)
		result   []*models.DateHistory
	)

	getOrder, err := c.OrderGetList(&models.OrderGetListRequest{})
	if err != nil {
		return nil, err
	}

	for _, value := range getOrder.Orders {
		if value.Status == true {
			toptimes[value.DateTime] += value.SumCount
		}
	}

	for k, v := range toptimes {
		result = append(result, &models.DateHistory{
			Date:  k,
			Count: v,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})

	return result, nil
}

// 10-task
func (c *Controller) ActiveUser() (string, error) {
	users := make(map[string]int)
	getorder, err := c.OrderGetList(&models.OrderGetListRequest{})
	if err != nil {
		return "", err
	}
	for _, value := range getorder.Orders {
		if value.Status == true {
			getproduct, err := c.GetByIdPoduct(&models.ProductPrimaryKey{Id: value.Id})
			if err != nil {
				return "", err
			}
			users[value.UserId] += value.SumCount * getproduct.Price
		}
	}
	user, sum := "", 0
	for key, value := range users {
		if sum < value {
			user = key
			sum = value
		}
	}
	getuser, err := c.UserGetById(&models.UserPrimaryKey{
		Id: user,
	})
	if err != nil {
		return "", err
	}
	return getuser.FirstName, nil
}
