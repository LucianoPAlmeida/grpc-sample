//
//  ViewController.swift
//  grpc-client
//
//  Created by Luciano Almeida on 01/06/19.
//  Copyright © 2019 Luciano Almeida. All rights reserved.
//

import UIKit

class ViewController: UIViewController {

    @IBOutlet weak var txtNumber: UITextField!
    @IBOutlet weak var lblName: UILabel!
    
    override func viewDidLoad() {
        super.viewDidLoad()

    }

    @IBAction func actionGet(_ sender: Any) {
        guard let number = Int(txtNumber.text ?? "") else { return }
        AppDriverService.findDriver(with: number) { [unowned self](result) in
            DispatchQueue.main.async {
                switch result {
                case .failure:
                    self.lblName.text = "Error :("
                case .success(let driver):
                    if let driver = driver {
                        self.lblName.text = "\(driver.name) \(number)"
                    } else {
                        self.lblName.text = "Driver \(number) not found on the server"
                    }
                }
            }
        }
    }

}

