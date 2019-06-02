//
//  DriverService.swift
//  grpc-client
//
//  Created by Luciano Almeida on 01/06/19.
//  Copyright © 2019 Luciano Almeida. All rights reserved.
//
import SwiftGRPC

class AppDriverService {
    static let grpcClient = DriverServiceClient(address: "localhost:50051", secure: false)
    
    class func findDriver(with number: Int, completion: @escaping (Result<Driver?, Error>) -> Void) {
        var request = DriverRequest()
        request.number = Int32(number)
        do {
            try grpcClient.findDriver(request, completion: { (response, result) in
                if result.success {
                    guard let driver = response else {
                        completion(.success(nil))
                        return
                    }
                    completion(.success(Driver(name: driver.name)))
                }
            })
        } catch let error {
            completion(.failure(error))
        }
        
    }
}
