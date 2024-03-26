# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BooksGet**](BooksApi.md#BooksGet) | **Get** /books | Get all books
[**BooksIdDelete**](BooksApi.md#BooksIdDelete) | **Delete** /books/{id} | Delete a book
[**BooksIdGet**](BooksApi.md#BooksIdGet) | **Get** /books/{id} | Get a book by ID
[**BooksIdPut**](BooksApi.md#BooksIdPut) | **Put** /books/{id} | Update a book
[**BooksPost**](BooksApi.md#BooksPost) | **Post** /books | Create a new book

# **BooksGet**
> []ModelsBook BooksGet(ctx, )
Get all books

Retrieve a list of all books

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ModelsBook**](models.Book.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BooksIdDelete**
> BooksIdDelete(ctx, id)
Delete a book

Delete a book by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Book ID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BooksIdGet**
> ModelsBook BooksIdGet(ctx, id)
Get a book by ID

Retrieve a book by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Book ID | 

### Return type

[**ModelsBook**](models.Book.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BooksIdPut**
> ModelsBook BooksIdPut(ctx, body, id)
Update a book

Update a book with the provided details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelsBook**](ModelsBook.md)| Book object | 
  **id** | **int32**| Book ID | 

### Return type

[**ModelsBook**](models.Book.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BooksPost**
> ModelsBook BooksPost(ctx, body)
Create a new book

Create a new book with the provided details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelsBook**](ModelsBook.md)| Book object | 

### Return type

[**ModelsBook**](models.Book.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

