/* tslint:disable */
/* eslint-disable */
/**
 * Goallery
 * This is the API Specification for Goallery.
 *
 * The version of the OpenAPI document: 1.23.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  AuthRequest,
  AuthResponse,
  ProblemDetails,
} from '../models/index';
import {
    AuthRequestFromJSON,
    AuthRequestToJSON,
    AuthResponseFromJSON,
    AuthResponseToJSON,
    ProblemDetailsFromJSON,
    ProblemDetailsToJSON,
} from '../models/index';

export interface GetTokenRequest {
    body: AuthRequest;
}

/**
 * 
 */
export class AuthApi extends runtime.BaseAPI {

    /**
     * Get JWT token
     * Get JWT token
     */
    async getTokenRaw(requestParameters: GetTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<AuthResponse>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling getToken.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/auth/login`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: AuthRequestToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => AuthResponseFromJSON(jsonValue));
    }

    /**
     * Get JWT token
     * Get JWT token
     */
    async getToken(requestParameters: GetTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<AuthResponse> {
        const response = await this.getTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
